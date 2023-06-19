// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	context "context"

	model "github.com/mattermost/mattermost/server/public/model"
	mock "github.com/stretchr/testify/mock"
)

// UploadSessionStore is an autogenerated mock type for the UploadSessionStore type
type UploadSessionStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *UploadSessionStore) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *UploadSessionStore) Get(ctx context.Context, id string) (*model.UploadSession, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.UploadSession
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.UploadSession, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.UploadSession); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UploadSession)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForUser provides a mock function with given fields: userID
func (_m *UploadSessionStore) GetForUser(userID string) ([]*model.UploadSession, error) {
	ret := _m.Called(userID)

	var r0 []*model.UploadSession
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*model.UploadSession, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(string) []*model.UploadSession); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.UploadSession)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: session
func (_m *UploadSessionStore) Save(session *model.UploadSession) (*model.UploadSession, error) {
	ret := _m.Called(session)

	var r0 *model.UploadSession
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.UploadSession) (*model.UploadSession, error)); ok {
		return rf(session)
	}
	if rf, ok := ret.Get(0).(func(*model.UploadSession) *model.UploadSession); ok {
		r0 = rf(session)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UploadSession)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.UploadSession) error); ok {
		r1 = rf(session)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: session
func (_m *UploadSessionStore) Update(session *model.UploadSession) error {
	ret := _m.Called(session)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.UploadSession) error); ok {
		r0 = rf(session)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUploadSessionStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewUploadSessionStore creates a new instance of UploadSessionStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUploadSessionStore(t mockConstructorTestingTNewUploadSessionStore) *UploadSessionStore {
	mock := &UploadSessionStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
