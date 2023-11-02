// Code generated by mockery v2.23.2. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	request "github.com/mattermost/mattermost/server/public/shared/request"
	mock "github.com/stretchr/testify/mock"
)

// LdapInterface is an autogenerated mock type for the LdapInterface type
type LdapInterface struct {
	mock.Mock
}

// CheckPassword provides a mock function with given fields: c, id, password
func (_m *LdapInterface) CheckPassword(c request.CTX, id string, password string) *model.AppError {
	ret := _m.Called(c, id, password)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, string, string) *model.AppError); ok {
		r0 = rf(c, id, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// CheckPasswordAuthData provides a mock function with given fields: c, authData, password
func (_m *LdapInterface) CheckPasswordAuthData(c request.CTX, authData string, password string) *model.AppError {
	ret := _m.Called(c, authData, password)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, string, string) *model.AppError); ok {
		r0 = rf(c, authData, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// CheckProviderAttributes provides a mock function with given fields: c, LS, ouser, patch
func (_m *LdapInterface) CheckProviderAttributes(c request.CTX, LS *model.LdapSettings, ouser *model.User, patch *model.UserPatch) string {
	ret := _m.Called(c, LS, ouser, patch)

	var r0 string
	if rf, ok := ret.Get(0).(func(request.CTX, *model.LdapSettings, *model.User, *model.UserPatch) string); ok {
		r0 = rf(c, LS, ouser, patch)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// DoLogin provides a mock function with given fields: c, id, password
func (_m *LdapInterface) DoLogin(c request.CTX, id string, password string) (*model.User, *model.AppError) {
	ret := _m.Called(c, id, password)

	var r0 *model.User
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, string, string) (*model.User, *model.AppError)); ok {
		return rf(c, id, password)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, string, string) *model.User); ok {
		r0 = rf(c, id, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, string, string) *model.AppError); ok {
		r1 = rf(c, id, password)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// FirstLoginSync provides a mock function with given fields: c, user, userAuthService, userAuthData, email
func (_m *LdapInterface) FirstLoginSync(c request.CTX, user *model.User, userAuthService string, userAuthData string, email string) *model.AppError {
	ret := _m.Called(c, user, userAuthService, userAuthData, email)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, *model.User, string, string, string) *model.AppError); ok {
		r0 = rf(c, user, userAuthService, userAuthData, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// GetADLdapIdFromSAMLId provides a mock function with given fields: c, authData
func (_m *LdapInterface) GetADLdapIdFromSAMLId(c request.CTX, authData string) string {
	ret := _m.Called(c, authData)

	var r0 string
	if rf, ok := ret.Get(0).(func(request.CTX, string) string); ok {
		r0 = rf(c, authData)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetAllGroupsPage provides a mock function with given fields: page, perPage, opts
func (_m *LdapInterface) GetAllGroupsPage(page int, perPage int, opts model.LdapGroupSearchOpts) ([]*model.Group, int, *model.AppError) {
	ret := _m.Called(page, perPage, opts)

	var r0 []*model.Group
	var r1 int
	var r2 *model.AppError
	if rf, ok := ret.Get(0).(func(int, int, model.LdapGroupSearchOpts) ([]*model.Group, int, *model.AppError)); ok {
		return rf(page, perPage, opts)
	}
	if rf, ok := ret.Get(0).(func(int, int, model.LdapGroupSearchOpts) []*model.Group); ok {
		r0 = rf(page, perPage, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, model.LdapGroupSearchOpts) int); ok {
		r1 = rf(page, perPage, opts)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(int, int, model.LdapGroupSearchOpts) *model.AppError); ok {
		r2 = rf(page, perPage, opts)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*model.AppError)
		}
	}

	return r0, r1, r2
}

// GetAllLdapUsers provides a mock function with given fields: c
func (_m *LdapInterface) GetAllLdapUsers(c request.CTX) ([]*model.User, *model.AppError) {
	ret := _m.Called(c)

	var r0 []*model.User
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX) ([]*model.User, *model.AppError)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(request.CTX) []*model.User); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX) *model.AppError); ok {
		r1 = rf(c)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetGroup provides a mock function with given fields: groupUID
func (_m *LdapInterface) GetGroup(groupUID string) (*model.Group, *model.AppError) {
	ret := _m.Called(groupUID)

	var r0 *model.Group
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string) (*model.Group, *model.AppError)); ok {
		return rf(groupUID)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Group); ok {
		r0 = rf(groupUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(groupUID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetSAMLIdFromADLdapId provides a mock function with given fields: c, authData
func (_m *LdapInterface) GetSAMLIdFromADLdapId(c request.CTX, authData string) string {
	ret := _m.Called(c, authData)

	var r0 string
	if rf, ok := ret.Get(0).(func(request.CTX, string) string); ok {
		r0 = rf(c, authData)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetUser provides a mock function with given fields: c, id
func (_m *LdapInterface) GetUser(c request.CTX, id string) (*model.User, *model.AppError) {
	ret := _m.Called(c, id)

	var r0 *model.User
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, string) (*model.User, *model.AppError)); ok {
		return rf(c, id)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, string) *model.User); ok {
		r0 = rf(c, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, string) *model.AppError); ok {
		r1 = rf(c, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetUserAttributes provides a mock function with given fields: id, attributes
func (_m *LdapInterface) GetUserAttributes(id string, attributes []string) (map[string]string, *model.AppError) {
	ret := _m.Called(id, attributes)

	var r0 map[string]string
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string, []string) (map[string]string, *model.AppError)); ok {
		return rf(id, attributes)
	}
	if rf, ok := ret.Get(0).(func(string, []string) map[string]string); ok {
		r0 = rf(id, attributes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []string) *model.AppError); ok {
		r1 = rf(id, attributes)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetVendorNameAndVendorVersion provides a mock function with given fields:
func (_m *LdapInterface) GetVendorNameAndVendorVersion() (string, string) {
	ret := _m.Called()

	var r0 string
	var r1 string
	if rf, ok := ret.Get(0).(func() (string, string)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// MigrateIDAttribute provides a mock function with given fields: c, toAttribute
func (_m *LdapInterface) MigrateIDAttribute(c request.CTX, toAttribute string) error {
	ret := _m.Called(c, toAttribute)

	var r0 error
	if rf, ok := ret.Get(0).(func(request.CTX, string) error); ok {
		r0 = rf(c, toAttribute)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RunTest provides a mock function with given fields:
func (_m *LdapInterface) RunTest() *model.AppError {
	ret := _m.Called()

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func() *model.AppError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// StartSynchronizeJob provides a mock function with given fields: c, waitForJobToFinish, includeRemovedMembers
func (_m *LdapInterface) StartSynchronizeJob(c request.CTX, waitForJobToFinish bool, includeRemovedMembers bool) (*model.Job, *model.AppError) {
	ret := _m.Called(c, waitForJobToFinish, includeRemovedMembers)

	var r0 *model.Job
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, bool, bool) (*model.Job, *model.AppError)); ok {
		return rf(c, waitForJobToFinish, includeRemovedMembers)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, bool, bool) *model.Job); ok {
		r0 = rf(c, waitForJobToFinish, includeRemovedMembers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Job)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CTX, bool, bool) *model.AppError); ok {
		r1 = rf(c, waitForJobToFinish, includeRemovedMembers)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// SwitchToLdap provides a mock function with given fields: c, userID, ldapID, ldapPassword
func (_m *LdapInterface) SwitchToLdap(c request.CTX, userID string, ldapID string, ldapPassword string) *model.AppError {
	ret := _m.Called(c, userID, ldapID, ldapPassword)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, string, string, string) *model.AppError); ok {
		r0 = rf(c, userID, ldapID, ldapPassword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// UpdateProfilePictureIfNecessary provides a mock function with given fields: _a0, _a1, _a2
func (_m *LdapInterface) UpdateProfilePictureIfNecessary(_a0 request.CTX, _a1 model.User, _a2 model.Session) {
	_m.Called(_a0, _a1, _a2)
}

type mockConstructorTestingTNewLdapInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewLdapInterface creates a new instance of LdapInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLdapInterface(t mockConstructorTestingTNewLdapInterface) *LdapInterface {
	mock := &LdapInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
