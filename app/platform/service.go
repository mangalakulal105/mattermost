// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package platform

import (
	"sync"
	"sync/atomic"

	"github.com/mattermost/mattermost-server/v6/app/featureflag"
	"github.com/mattermost/mattermost-server/v6/config"
	"github.com/mattermost/mattermost-server/v6/einterfaces"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
)

// PlatformService is the service for the platform related tasks. It is
// responsible for non-entity related functionalities that are required
// by a product such as database access, configuration access, licensing etc.
type PlatformService struct {
	serviceConfig ServiceConfig
	configStore   *config.Store

	metrics *platformMetrics

	featureFlagSynchronizerMutex sync.Mutex
	featureFlagSynchronizer      *featureflag.Synchronizer
	featureFlagStop              chan struct{}
	featureFlagStopped           chan struct{}

	licenseValue atomic.Value

	cluster einterfaces.ClusterInterface
}

// New creates a new PlatformService.
func New(sc ServiceConfig) (*PlatformService, error) {
	if err := sc.validate(); err != nil {
		return nil, err
	}

	ps := &PlatformService{
		serviceConfig: sc,
		configStore:   sc.ConfigStore,
		cluster:       sc.Cluster,
	}

	ps.metrics = newPlatformMetrics(sc.Metrics, ps.configStore.Get)

	return ps, nil
}

func (ps *PlatformService) ShutdownMetrics() {
	if ps.metrics != nil {
		ps.metrics.stopMetricsServer()
	}
}

func (ps *PlatformService) ShutdownConfig() {
	if ps.configStore != nil {
		err := ps.configStore.Close()
		if err != nil {
			mlog.Warn("Failed to close config store", mlog.Err(err))
		}
	}
}
