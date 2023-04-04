// Code generated by mockery v2.10.4. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/v6/model"
	mock "github.com/stretchr/testify/mock"
)

// PostPersistentNotificationStore is an autogenerated mock type for the PostPersistentNotificationStore type
type PostPersistentNotificationStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: postIds
func (_m *PostPersistentNotificationStore) Delete(postIds []string) error {
	ret := _m.Called(postIds)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(postIds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByChannel provides a mock function with given fields: channelIds
func (_m *PostPersistentNotificationStore) DeleteByChannel(channelIds []string) error {
	ret := _m.Called(channelIds)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(channelIds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByTeam provides a mock function with given fields: teamIds
func (_m *PostPersistentNotificationStore) DeleteByTeam(teamIds []string) error {
	ret := _m.Called(teamIds)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(teamIds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteExpired provides a mock function with given fields: maxSentCount
func (_m *PostPersistentNotificationStore) DeleteExpired(maxSentCount int16) error {
	ret := _m.Called(maxSentCount)

	var r0 error
	if rf, ok := ret.Get(0).(func(int16) error); ok {
		r0 = rf(maxSentCount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: params
func (_m *PostPersistentNotificationStore) Get(params model.GetPersistentNotificationsPostsParams) ([]*model.PostPersistentNotifications, error) {
	ret := _m.Called(params)

	var r0 []*model.PostPersistentNotifications
	if rf, ok := ret.Get(0).(func(model.GetPersistentNotificationsPostsParams) []*model.PostPersistentNotifications); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.PostPersistentNotifications)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.GetPersistentNotificationsPostsParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSingle provides a mock function with given fields: postID
func (_m *PostPersistentNotificationStore) GetSingle(postID string) (*model.PostPersistentNotifications, error) {
	ret := _m.Called(postID)

	var r0 *model.PostPersistentNotifications
	if rf, ok := ret.Get(0).(func(string) *model.PostPersistentNotifications); ok {
		r0 = rf(postID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PostPersistentNotifications)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(postID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateLastActivity provides a mock function with given fields: postIds
func (_m *PostPersistentNotificationStore) UpdateLastActivity(postIds []string) error {
	ret := _m.Called(postIds)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(postIds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
