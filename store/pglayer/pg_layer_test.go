// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package pglayer

import (
	"testing"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/store"
	"github.com/mattermost/mattermost-server/v5/store/sqlstore"
	"github.com/mattermost/mattermost-server/v5/store/storetest"
)

var pgStore *PgLayer

func StoreTest(t *testing.T, f func(*testing.T, store.Store)) {
	defer func() {
		if err := recover(); err != nil {
			tearDownStores()
			panic(err)
		}
	}()
	t.Run("PG", func(t *testing.T) { f(t, pgStore) })
}

func StoreTestWithSqlSupplier(t *testing.T, f func(*testing.T, store.Store, storetest.SqlSupplier)) {
	defer func() {
		if err := recover(); err != nil {
			tearDownStores()
			panic(err)
		}
	}()
	t.Run("PG", func(t *testing.T) { f(t, pgStore, pgStore) })
}

func initStores() {
	defer func() {
		if err := recover(); err != nil {
			tearDownStores()
			panic(err)
		}
	}()
	pgStore = NewPgLayer(*sqlstore.NewSqlSupplier(*storetest.MakeSqlSettings(model.DATABASE_DRIVER_POSTGRES), nil))
	pgStore.Init()
	pgStore.DropAllTables()
	pgStore.MarkSystemRanUnitTests()
}

func tearDownStores() {
	pgStore.Close()
}
