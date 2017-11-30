// Code generated by mockery v1.0.0

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import store "github.com/mattermost/mattermost-server/store"

// ChannelMemberHistoryStore is an autogenerated mock type for the ChannelMemberHistoryStore type
type ChannelMemberHistoryStore struct {
	mock.Mock
}

// GetUsersInChannelDuring provides a mock function with given fields: startTime, endTime, channelId
func (_m *ChannelMemberHistoryStore) GetUsersInChannelDuring(startTime int64, endTime int64, channelId string) store.StoreChannel {
	ret := _m.Called(startTime, endTime, channelId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int64, int64, string) store.StoreChannel); ok {
		r0 = rf(startTime, endTime, channelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// LogJoinEvent provides a mock function with given fields: userId, channelId, joinTime
func (_m *ChannelMemberHistoryStore) LogJoinEvent(userId string, channelId string, joinTime int64) store.StoreChannel {
	ret := _m.Called(userId, channelId, joinTime)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, int64) store.StoreChannel); ok {
		r0 = rf(userId, channelId, joinTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// LogLeaveEvent provides a mock function with given fields: userId, channelId, leaveTime
func (_m *ChannelMemberHistoryStore) LogLeaveEvent(userId string, channelId string, leaveTime int64) store.StoreChannel {
	ret := _m.Called(userId, channelId, leaveTime)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, int64) store.StoreChannel); ok {
		r0 = rf(userId, channelId, leaveTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PurgeHistoryBefore provides a mock function with given fields: time, channelId
func (_m *ChannelMemberHistoryStore) PurgeHistoryBefore(time int64, channelId string) store.StoreChannel {
	ret := _m.Called(time, channelId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int64, string) store.StoreChannel); ok {
		r0 = rf(time, channelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
