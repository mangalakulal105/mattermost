// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package platform

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/mattermost/mattermost-server/v6/config"
	"github.com/mattermost/mattermost-server/v6/einterfaces"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/product"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
)

// ServiceConfig is used to initialize the PlatformService.
// The mandatory fields will be checked during the initialization of the service.
type ServiceConfig struct {
	// Mandatory fields
	ConfigStore  *config.Store
	StartMetrics bool // TODO: find an elegant way to start/stop metrics server by default
	// Optional fields
	Metrics einterfaces.MetricsInterface
	Cluster einterfaces.ClusterInterface
}

func (c *ServiceConfig) validate() error {
	// Mandatory fields need to be checked here
	if c.ConfigStore == nil {
		return errors.New("ConfigStore is required")
	}

	return nil
}

// ensure the config wrapper implements `product.ConfigService`
var _ product.ConfigService = (*PlatformService)(nil)

func (ps *PlatformService) Config() *model.Config {
	return ps.configStore.Get()
}

func (ps *PlatformService) AddConfigListener(listener func(*model.Config, *model.Config)) string {
	return ps.configStore.AddListener(listener)
}

func (ps *PlatformService) RemoveConfigListener(id string) {
	ps.configStore.RemoveListener(id)
}

func (ps *PlatformService) UpdateConfig(f func(*model.Config)) {
	if ps.configStore.IsReadOnly() {
		return
	}
	old := ps.Config()
	updated := old.Clone()
	f(updated)
	if _, _, err := ps.configStore.Set(updated); err != nil {
		mlog.Error("Failed to update config", mlog.Err(err))
	}
}

// SaveConfig replaces the active configuration, optionally notifying cluster peers.
// It returns both the previous and current configs.
func (ps *PlatformService) SaveConfig(newCfg *model.Config, sendConfigChangeClusterMessage bool) (*model.Config, *model.Config, *model.AppError) {
	oldCfg, newCfg, err := ps.configStore.Set(newCfg)
	if errors.Is(err, config.ErrReadOnlyConfiguration) {
		return nil, nil, model.NewAppError("saveConfig", "ent.cluster.save_config.error", nil, err.Error(), http.StatusForbidden)
	} else if err != nil {
		return nil, nil, model.NewAppError("saveConfig", "app.save_config.app_error", nil, err.Error(), http.StatusInternalServerError)
	}

	if ps.serviceConfig.StartMetrics && *ps.Config().MetricsSettings.Enable {
		if ps.metrics.metricsImpl != nil {
			ps.metrics.metricsImpl.Register()
		}
		ps.RestartMetrics()
	} else {
		ps.ShutdownMetrics()
	}

	if ps.cluster != nil {
		err := ps.cluster.ConfigChanged(ps.configStore.RemoveEnvironmentOverrides(oldCfg),
			ps.configStore.RemoveEnvironmentOverrides(newCfg), sendConfigChangeClusterMessage)
		if err != nil {
			return nil, nil, err
		}
	}

	return oldCfg, newCfg, nil
}

func (ps *PlatformService) ReloadConfig() error {
	if err := ps.configStore.Load(); err != nil {
		return err
	}
	return nil
}

func (ps *PlatformService) GetEnvironmentOverridesWithFilter(filter func(reflect.StructField) bool) map[string]interface{} {
	return ps.configStore.GetEnvironmentOverridesWithFilter(filter)
}

func (ps *PlatformService) GetEnvironmentOverrides() map[string]interface{} {
	return ps.configStore.GetEnvironmentOverrides()
}

func (ps *PlatformService) DescribeConfig() string {
	return ps.configStore.String()
}

func (ps *PlatformService) CleanUpConfig() error {
	return ps.configStore.CleanUp()
}

// ConfigureLogger applies the specified configuration to a logger.
func (ps *PlatformService) ConfigureLogger(name string, logger *mlog.Logger, logSettings *model.LogSettings, getPath func(string) string) error {
	// Advanced logging is E20 only, however logging must be initialized before the license
	// file is loaded.  If no valid E20 license exists then advanced logging will be
	// shutdown once license is loaded/checked.
	var err error
	dsn := *logSettings.AdvancedLoggingConfig
	var logConfigSrc config.LogConfigSrc
	if dsn != "" {
		logConfigSrc, err = config.NewLogConfigSrc(dsn, ps.configStore)
		if err != nil {
			return fmt.Errorf("invalid config source for %s, %w", name, err)
		}
		mlog.Info("Loaded configuration for "+name, mlog.String("source", dsn))
	}

	cfg, err := config.MloggerConfigFromLoggerConfig(logSettings, logConfigSrc, getPath)
	if err != nil {
		return fmt.Errorf("invalid config source for %s, %w", name, err)
	}

	if err := logger.ConfigureTargets(cfg, nil); err != nil {
		return fmt.Errorf("invalid config for %s, %w", name, err)
	}
	return nil
}

func (ps *PlatformService) GetConfigStore() *config.Store {
	return ps.configStore
}
