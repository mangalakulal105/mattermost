// Code generated by mockery v2.42.2. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	request "github.com/mattermost/mattermost/server/public/shared/request"
	mock "github.com/stretchr/testify/mock"
)

// SessionStore is an autogenerated mock type for the SessionStore type
type SessionStore struct {
	mock.Mock
}

// AnalyticsSessionCount provides a mock function with given fields:
func (_m *SessionStore) AnalyticsSessionCount() (int64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AnalyticsSessionCount")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Cleanup provides a mock function with given fields: expiryTime, batchSize
func (_m *SessionStore) Cleanup(expiryTime int64, batchSize int64) error {
	ret := _m.Called(expiryTime, batchSize)

	if len(ret) == 0 {
		panic("no return value specified for Cleanup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, int64) error); ok {
		r0 = rf(expiryTime, batchSize)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: c, sessionIDOrToken
func (_m *SessionStore) Get(c request.CTX, sessionIDOrToken string) (*model.Session, error) {
	ret := _m.Called(c, sessionIDOrToken)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *model.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(request.CTX, string) (*model.Session, error)); ok {
		return rf(c, sessionIDOrToken)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, string) *model.Session); ok {
		r0 = rf(c, sessionIDOrToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, string) error); ok {
		r1 = rf(c, sessionIDOrToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLRUSessions provides a mock function with given fields: c, userID, limit, offset
func (_m *SessionStore) GetLRUSessions(c request.CTX, userID string, limit uint64, offset uint64) ([]*model.Session, error) {
	ret := _m.Called(c, userID, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetLRUSessions")
	}

	var r0 []*model.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(request.CTX, string, uint64, uint64) ([]*model.Session, error)); ok {
		return rf(c, userID, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, string, uint64, uint64) []*model.Session); ok {
		r0 = rf(c, userID, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, string, uint64, uint64) error); ok {
		r1 = rf(c, userID, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSessions provides a mock function with given fields: c, userID
func (_m *SessionStore) GetSessions(c request.CTX, userID string) ([]*model.Session, error) {
	ret := _m.Called(c, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetSessions")
	}

	var r0 []*model.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(request.CTX, string) ([]*model.Session, error)); ok {
		return rf(c, userID)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, string) []*model.Session); ok {
		r0 = rf(c, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, string) error); ok {
		r1 = rf(c, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSessionsExpired provides a mock function with given fields: thresholdMillis, mobileOnly, unnotifiedOnly
func (_m *SessionStore) GetSessionsExpired(thresholdMillis int64, mobileOnly bool, unnotifiedOnly bool) ([]*model.Session, error) {
	ret := _m.Called(thresholdMillis, mobileOnly, unnotifiedOnly)

	if len(ret) == 0 {
		panic("no return value specified for GetSessionsExpired")
	}

	var r0 []*model.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, bool, bool) ([]*model.Session, error)); ok {
		return rf(thresholdMillis, mobileOnly, unnotifiedOnly)
	}
	if rf, ok := ret.Get(0).(func(int64, bool, bool) []*model.Session); ok {
		r0 = rf(thresholdMillis, mobileOnly, unnotifiedOnly)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, bool, bool) error); ok {
		r1 = rf(thresholdMillis, mobileOnly, unnotifiedOnly)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSessionsWithActiveDeviceIds provides a mock function with given fields: userID
func (_m *SessionStore) GetSessionsWithActiveDeviceIds(userID string) ([]*model.Session, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetSessionsWithActiveDeviceIds")
	}

	var r0 []*model.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*model.Session, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(string) []*model.Session); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentDeleteSessionsByUser provides a mock function with given fields: teamID
func (_m *SessionStore) PermanentDeleteSessionsByUser(teamID string) error {
	ret := _m.Called(teamID)

	if len(ret) == 0 {
		panic("no return value specified for PermanentDeleteSessionsByUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(teamID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Remove provides a mock function with given fields: sessionIDOrToken
func (_m *SessionStore) Remove(sessionIDOrToken string) error {
	ret := _m.Called(sessionIDOrToken)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(sessionIDOrToken)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveAllSessions provides a mock function with given fields:
func (_m *SessionStore) RemoveAllSessions() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RemoveAllSessions")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: c, session
func (_m *SessionStore) Save(c request.CTX, session *model.Session) (*model.Session, error) {
	ret := _m.Called(c, session)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 *model.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Session) (*model.Session, error)); ok {
		return rf(c, session)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, *model.Session) *model.Session); ok {
		r0 = rf(c, session)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, *model.Session) error); ok {
		r1 = rf(c, session)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDeviceId provides a mock function with given fields: id, deviceID, expiresAt
func (_m *SessionStore) UpdateDeviceId(id string, deviceID string, expiresAt int64) (string, error) {
	ret := _m.Called(id, deviceID, expiresAt)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDeviceId")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int64) (string, error)); ok {
		return rf(id, deviceID, expiresAt)
	}
	if rf, ok := ret.Get(0).(func(string, string, int64) string); ok {
		r0 = rf(id, deviceID, expiresAt)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, int64) error); ok {
		r1 = rf(id, deviceID, expiresAt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateExpiredNotify provides a mock function with given fields: sessionid, notified
func (_m *SessionStore) UpdateExpiredNotify(sessionid string, notified bool) error {
	ret := _m.Called(sessionid, notified)

	if len(ret) == 0 {
		panic("no return value specified for UpdateExpiredNotify")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool) error); ok {
		r0 = rf(sessionid, notified)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateExpiresAt provides a mock function with given fields: sessionID, timestamp
func (_m *SessionStore) UpdateExpiresAt(sessionID string, timestamp int64) error {
	ret := _m.Called(sessionID, timestamp)

	if len(ret) == 0 {
		panic("no return value specified for UpdateExpiresAt")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64) error); ok {
		r0 = rf(sessionID, timestamp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateLastActivityAt provides a mock function with given fields: sessionID, timestamp
func (_m *SessionStore) UpdateLastActivityAt(sessionID string, timestamp int64) error {
	ret := _m.Called(sessionID, timestamp)

	if len(ret) == 0 {
		panic("no return value specified for UpdateLastActivityAt")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64) error); ok {
		r0 = rf(sessionID, timestamp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProps provides a mock function with given fields: session
func (_m *SessionStore) UpdateProps(session *model.Session) error {
	ret := _m.Called(session)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProps")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Session) error); ok {
		r0 = rf(session)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRoles provides a mock function with given fields: userID, roles
func (_m *SessionStore) UpdateRoles(userID string, roles string) (string, error) {
	ret := _m.Called(userID, roles)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRoles")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(userID, roles)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(userID, roles)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, roles)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSessionStore creates a new instance of SessionStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSessionStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *SessionStore {
	mock := &SessionStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
