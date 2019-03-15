// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package config

import (
	"bytes"
	"io"
	"reflect"
	"sync"

	"github.com/mattermost/mattermost-server/model"
	"github.com/pkg/errors"
)

// commonStore enables code sharing between different backing implementations
type commonStore struct {
	emitter

	configLock             sync.RWMutex
	config                 *model.Config
	configWithoutOverrides *model.Config
	environmentOverrides   map[string]interface{}
}

// Get fetches the current, cached configuration.
func (cs *commonStore) Get() *model.Config {
	cs.configLock.RLock()
	defer cs.configLock.RUnlock()

	return cs.config
}

// GetEnvironmentOverrides fetches the configuration fields overridden by environment variables.
func (cs *commonStore) GetEnvironmentOverrides() map[string]interface{} {
	cs.configLock.RLock()
	defer cs.configLock.RUnlock()

	return cs.environmentOverrides
}

// set replaces the current configuration in its entirety, and updates the backing store
// using the persist function argument.
//
// This function assumes no lock has been acquired, as it acquires a write lock itself.
func (cs *commonStore) set(newCfg *model.Config, validate func(*model.Config) error, persist func(*model.Config) error) (*model.Config, error) {
	cs.configLock.Lock()
	var unlockOnce sync.Once
	defer unlockOnce.Do(cs.configLock.Unlock)

	oldCfg := cs.config

	// TODO: disallow attempting to save a directly modified config (comparing pointers). This
	// wouldn't be an exhaustive check, given the use of pointers throughout the data
	// structure, but might prevent common mistakes. Requires upstream changes first.
	// if newCfg == oldCfg {
	// 	return nil, errors.New("old configuration modified instead of cloning")
	// }

	newCfg = newCfg.Clone()
	newCfg.SetDefaults()

	// Sometimes the config is received with "fake" data in sensitive fields. Apply the real
	// data from the existing config as necessary.
	desanitize(oldCfg, newCfg)

	if validate != nil {
		if err := validate(newCfg); err != nil {
			return nil, errors.Wrap(err, "new configuration is invalid")
		}
	}

	//if err := persist(cs.removeEnvOverrides(newCfg)); err != nil {
	if err := persist(cs.removeEnvOverrides(newCfg)); err != nil {
		return nil, errors.Wrap(err, "failed to persist")
	}

	cs.config = newCfg

	unlockOnce.Do(cs.configLock.Unlock)

	// Notify listeners synchronously. Ideally, this would be asynchronous, but existing code
	// assumes this and there would be increased complexity to avoid racing updates.
	cs.invokeConfigListeners(oldCfg, newCfg)

	return oldCfg, nil
}

// load updates the current configuration from the given io.ReadCloser.
//
// This function assumes no lock has been acquired, as it acquires a write lock itself.
func (cs *commonStore) load(f io.ReadCloser, needsSave bool, validate func(*model.Config) error, persist func(*model.Config) error) error {
	// Split f into two so that we can have a configuration that does not have environment overrides
	f2 := new(bytes.Buffer)
	tee := io.TeeReader(f, f2)

	allowEnvironmentOverrides := true
	loadedCfg, environmentOverrides, err := unmarshalConfig(tee, allowEnvironmentOverrides)
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal config with env overrides")
	}

	// Keep track of the original values that the Environment settings overrode
	loadedCfgWithoutEnvOverrides, _, err := unmarshalConfig(f2, false)
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal config without env overrides")
	}

	// SetDefaults generates various keys and salts if not previously configured. Determine if
	// such a change will be made before invoking.
	needsSave = needsSave || loadedCfg.SqlSettings.AtRestEncryptKey == nil || len(*loadedCfg.SqlSettings.AtRestEncryptKey) == 0
	needsSave = needsSave || loadedCfg.FileSettings.PublicLinkSalt == nil || len(*loadedCfg.FileSettings.PublicLinkSalt) == 0
	needsSave = needsSave || loadedCfg.EmailSettings.InviteSalt == nil || len(*loadedCfg.EmailSettings.InviteSalt) == 0

	loadedCfg.SetDefaults()

	if validate != nil {
		if err = validate(loadedCfg); err != nil {
			return errors.Wrap(err, "invalid config")
		}
	}

	if changed := fixConfig(loadedCfg); changed {
		needsSave = true
	}

	cs.configLock.Lock()
	var unlockOnce sync.Once
	defer unlockOnce.Do(cs.configLock.Unlock)

	if needsSave && persist != nil {
		cfgWithoutEnvOverrides := removeEnvOverrides(loadedCfg, loadedCfgWithoutEnvOverrides, environmentOverrides)
		if err = persist(cfgWithoutEnvOverrides); err != nil {
			return errors.Wrap(err, "failed to persist required changes after load")
		}
	}

	oldCfg := cs.config
	cs.config = loadedCfg
	cs.configWithoutOverrides = loadedCfgWithoutEnvOverrides
	cs.environmentOverrides = environmentOverrides

	unlockOnce.Do(cs.configLock.Unlock)

	// Notify listeners synchronously. Ideally, this would be asynchronous, but existing code
	// assumes this and there would be increased complexity to avoid racing updates.
	cs.invokeConfigListeners(oldCfg, loadedCfg)

	return nil
}

// validate checks if the given configuration is valid
func (cs *commonStore) validate(cfg *model.Config) error {
	if err := cfg.IsValid(); err != nil {
		return err
	}

	return nil
}

// removeEnvOverrides takes the newCfg provided and adds information stored in the commonStore,
// then delegates the task to the removeEnvOverrides function
func (cs *commonStore) removeEnvOverrides(newCfg *model.Config) *model.Config {
	return removeEnvOverrides(newCfg, cs.configWithoutOverrides, cs.environmentOverrides)
}

// removeEnvOverrides returns a new config without the current environment overrides.
// If a config variable has an environment override, that variable is set to the value that was
// read from the store.
func removeEnvOverrides(cfg, cfgWithoutEnv *model.Config, envOverrides map[string]interface{}) *model.Config {
	paths := getPaths(envOverrides)
	newCfg := cfg.Clone()
	for _, path := range paths {
		currentVal := getVal(newCfg, path)
		originalVal := getVal(cfgWithoutEnv, path)
		if currentVal.Interface() != originalVal.Interface() {
			newVal := getVal(newCfg, path)
			newVal.Set(originalVal)
		}
	}
	return newCfg
}

// getPaths turns a nested map into a slice of paths describing the keys of the map. Eg:
// map[string]map[string]map[string]bool{"this":{"is first":{"path":true}, "is second":{"path":true}))) is turned into:
// [][]string{{"this", "is first", "path"}, {"this", "is second", "path"}}
func getPaths(m map[string]interface{}) [][]string {
	return getPathsRec(m, nil)
}

// getPathsRec assembles the paths (see `getPaths` above)
func getPathsRec(src interface{}, curPath []string) [][]string {
	if srcMap, ok := src.(map[string]interface{}); ok {
		paths := [][]string{}
		for k, v := range srcMap {
			paths = append(paths, getPathsRec(v, append(curPath, k))...)
		}
		return paths
	}

	return [][]string{curPath}
}

// getVal walks `src` (here it starts with a model.Config, then recurses into its leaves)
// and returns the reflect.Value of the leaf at the end `path`
func getVal(src interface{}, path []string) reflect.Value {
	var val reflect.Value
	if reflect.ValueOf(src).Kind() == reflect.Ptr {
		val = reflect.ValueOf(src).Elem().FieldByName(path[0])
	} else {
		val = reflect.ValueOf(src).FieldByName(path[0])
	}
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() == reflect.Struct {
		return getVal(val.Interface(), path[1:])
	}
	return val
}
