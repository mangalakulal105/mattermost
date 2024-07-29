// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/v8/einterfaces"
	"github.com/redis/rueidis"
)

// CacheOptions contains options for initializing a cache
type CacheOptions struct {
	Size                   int
	DefaultExpiry          time.Duration
	Name                   string
	InvalidateClusterEvent model.ClusterEvent
	Striped                bool
	// StripedBuckets is used only by LRUStriped and shouldn't be greater than the number
	// of CPUs available on the machine running this cache.
	StripedBuckets int
}

// Provider is a provider for Cache
type Provider interface {
	// NewCache creates a new cache with given options.
	NewCache(opts *CacheOptions) (Cache, error)
	// Connect opens a new connection to the cache using specific provider parameters.
	Connect() (string, error)
	// SetMetrics
	SetMetrics(metrics einterfaces.MetricsInterface)
	// Close releases any resources used by the cache provider.
	Close() error
	// Type returns what type of cache it generates.
	Type() string
}

type cacheProvider struct {
}

// NewProvider creates a new CacheProvider
func NewProvider() Provider {
	return &cacheProvider{}
}

// NewCache creates a new cache with given opts
func (c *cacheProvider) NewCache(opts *CacheOptions) (Cache, error) {
	if opts.Striped {
		return NewLRUStriped(opts)
	}
	return NewLRU(opts), nil
}

// Connect opens a new connection to the cache using specific provider parameters.
func (c *cacheProvider) Connect() (string, error) {
	return "OK", nil
}

func (c *cacheProvider) SetMetrics(metrics einterfaces.MetricsInterface) {
}

// Close releases any resources used by the cache provider.
func (c *cacheProvider) Close() error {
	return nil
}

func (c *cacheProvider) Type() string {
	return model.CacheTypeLRU
}

type redisProvider struct {
	client  rueidis.Client
	metrics einterfaces.MetricsInterface
}

type RedisOptions struct {
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

// NewProvider creates a new CacheProvider
func NewRedisProvider(opts *RedisOptions) (Provider, error) {
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress:       []string{opts.RedisAddr},
		Password:          opts.RedisPassword,
		SelectDB:          opts.RedisDB,
		ForceSingleClient: true,
		CacheSizeEachConn: 16 * (1 << 20), // 16MiB local cache size
		//TODO: look into MaxFlushDelay
	})
	if err != nil {
		return nil, err
	}
	return &redisProvider{client: client}, nil
}

// NewCache creates a new cache with given opts
func (r *redisProvider) NewCache(opts *CacheOptions) (Cache, error) {
	rr, err := NewRedis(opts, r.client)
	rr.metrics = r.metrics
	return rr, err
}

// Connect opens a new connection to the cache using specific provider parameters.
func (r *redisProvider) Connect() (string, error) {
	res, err := r.client.Do(context.Background(), r.client.B().Ping().Build()).ToString()
	if err != nil {
		return "", fmt.Errorf("unable to establish connection with redis: %v", err)
	}
	return res, nil
}

func (r *redisProvider) SetMetrics(metrics einterfaces.MetricsInterface) {
	r.metrics = metrics
}

func (r *redisProvider) Type() string {
	return model.CacheTypeRedis
}

// Close releases any resources used by the cache provider.
func (r *redisProvider) Close() error {
	r.client.Close()
	return nil
}
