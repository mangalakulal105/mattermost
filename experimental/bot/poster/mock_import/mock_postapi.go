// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-api/experimental/bot/poster (interfaces: PostAPI)

// Package mock_import is a generated GoMock package.
package mock_import

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/mattermost/mattermost-server/v5/model"
	reflect "reflect"
)

// MockPostAPI is a mock of PostAPI interface
type MockPostAPI struct {
	ctrl     *gomock.Controller
	recorder *MockPostAPIMockRecorder
}

// MockPostAPIMockRecorder is the mock recorder for MockPostAPI
type MockPostAPIMockRecorder struct {
	mock *MockPostAPI
}

// NewMockPostAPI creates a new mock instance
func NewMockPostAPI(ctrl *gomock.Controller) *MockPostAPI {
	mock := &MockPostAPI{ctrl: ctrl}
	mock.recorder = &MockPostAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostAPI) EXPECT() *MockPostAPIMockRecorder {
	return m.recorder
}

// DM mocks base method
func (m *MockPostAPI) DM(arg0, arg1 string, arg2 *model.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DM", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DM indicates an expected call of DM
func (mr *MockPostAPIMockRecorder) DM(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DM", reflect.TypeOf((*MockPostAPI)(nil).DM), arg0, arg1, arg2)
}

// DeletePost mocks base method
func (m *MockPostAPI) DeletePost(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost
func (mr *MockPostAPIMockRecorder) DeletePost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockPostAPI)(nil).DeletePost), arg0)
}

// GetPost mocks base method
func (m *MockPostAPI) GetPost(arg0 string) (*model.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", arg0)
	ret0, _ := ret[0].(*model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost
func (mr *MockPostAPIMockRecorder) GetPost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockPostAPI)(nil).GetPost), arg0)
}

// SendEphemeralPost mocks base method
func (m *MockPostAPI) SendEphemeralPost(arg0 string, arg1 *model.Post) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendEphemeralPost", arg0, arg1)
}

// SendEphemeralPost indicates an expected call of SendEphemeralPost
func (mr *MockPostAPIMockRecorder) SendEphemeralPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEphemeralPost", reflect.TypeOf((*MockPostAPI)(nil).SendEphemeralPost), arg0, arg1)
}

// UpdatePost mocks base method
func (m *MockPostAPI) UpdatePost(arg0 *model.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePost indicates an expected call of UpdatePost
func (mr *MockPostAPIMockRecorder) UpdatePost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockPostAPI)(nil).UpdatePost), arg0)
}
