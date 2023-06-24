// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	mock "github.com/stretchr/testify/mock"
)

// LdapSyncInterface is an autogenerated mock type for the LdapSyncInterface type
type LdapSyncInterface struct {
	mock.Mock
}

// MakeScheduler provides a mock function with given fields:
func (_m *LdapSyncInterface) MakeScheduler() model.Scheduler {
	ret := _m.Called()

	var r0 model.Scheduler
	if rf, ok := ret.Get(0).(func() model.Scheduler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Scheduler)
		}
	}

	return r0
}

// MakeWorker provides a mock function with given fields:
func (_m *LdapSyncInterface) MakeWorker() model.Worker {
	ret := _m.Called()

	var r0 model.Worker
	if rf, ok := ret.Get(0).(func() model.Worker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Worker)
		}
	}

	return r0
}

type mockConstructorTestingTNewLdapSyncInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewLdapSyncInterface creates a new instance of LdapSyncInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLdapSyncInterface(t mockConstructorTestingTNewLdapSyncInterface) *LdapSyncInterface {
	mock := &LdapSyncInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
