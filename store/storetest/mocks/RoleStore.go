// Code generated by mockery v2.10.4. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	context "context"

	model "github.com/mattermost/mattermost-server/v6/model"
	mock "github.com/stretchr/testify/mock"
)

// RoleStore is an autogenerated mock type for the RoleStore type
type RoleStore struct {
	mock.Mock
}

// AllChannelSchemeRoles provides a mock function with given fields:
func (_m *RoleStore) AllChannelSchemeRoles() ([]*model.Role, error) {
	ret := _m.Called()

	var r0 []*model.Role
	if rf, ok := ret.Get(0).(func() []*model.Role); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChannelHigherScopedPermissions provides a mock function with given fields: roleNames
func (_m *RoleStore) ChannelHigherScopedPermissions(roleNames []string) (map[string]*model.RolePermissions, error) {
	ret := _m.Called(roleNames)

	var r0 map[string]*model.RolePermissions
	if rf, ok := ret.Get(0).(func([]string) map[string]*model.RolePermissions); ok {
		r0 = rf(roleNames)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*model.RolePermissions)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(roleNames)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChannelRolesUnderTeamRole provides a mock function with given fields: roleName
func (_m *RoleStore) ChannelRolesUnderTeamRole(roleName string) ([]*model.Role, error) {
	ret := _m.Called(roleName)

	var r0 []*model.Role
	if rf, ok := ret.Get(0).(func(string) []*model.Role); ok {
		r0 = rf(roleName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(roleName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: roleID
func (_m *RoleStore) Delete(roleID string) (*model.Role, error) {
	ret := _m.Called(roleID)

	var r0 *model.Role
	if rf, ok := ret.Get(0).(func(string) *model.Role); ok {
		r0 = rf(roleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(roleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: roleID
func (_m *RoleStore) Get(roleID string) (*model.Role, error) {
	ret := _m.Called(roleID)

	var r0 *model.Role
	if rf, ok := ret.Get(0).(func(string) *model.Role); ok {
		r0 = rf(roleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(roleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *RoleStore) GetAll() ([]*model.Role, error) {
	ret := _m.Called()

	var r0 []*model.Role
	if rf, ok := ret.Get(0).(func() []*model.Role); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *RoleStore) GetByName(ctx context.Context, name string) (*model.Role, error) {
	ret := _m.Called(ctx, name)

	var r0 *model.Role
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Role); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByNames provides a mock function with given fields: names
func (_m *RoleStore) GetByNames(names []string) ([]*model.Role, error) {
	ret := _m.Called(names)

	var r0 []*model.Role
	if rf, ok := ret.Get(0).(func([]string) []*model.Role); ok {
		r0 = rf(names)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(names)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PermanentDeleteAll provides a mock function with given fields:
func (_m *RoleStore) PermanentDeleteAll() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: role
func (_m *RoleStore) Save(role *model.Role) (*model.Role, error) {
	ret := _m.Called(role)

	var r0 *model.Role
	if rf, ok := ret.Get(0).(func(*model.Role) *model.Role); ok {
		r0 = rf(role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Role) error); ok {
		r1 = rf(role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
