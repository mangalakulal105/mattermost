// Code generated by mockery v2.10.4. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/v6/model"
	mock "github.com/stretchr/testify/mock"
)

// TrueUpReviewStore is an autogenerated mock type for the TrueUpReviewStore type
type TrueUpReviewStore struct {
	mock.Mock
}

// CreateTrueUpReviewStatusRecord provides a mock function with given fields: reviewStatus
func (_m *TrueUpReviewStore) CreateTrueUpReviewStatusRecord(reviewStatus *model.TrueUpReviewStatus) (*model.TrueUpReviewStatus, error) {
	ret := _m.Called(reviewStatus)

	var r0 *model.TrueUpReviewStatus
	if rf, ok := ret.Get(0).(func(*model.TrueUpReviewStatus) *model.TrueUpReviewStatus); ok {
		r0 = rf(reviewStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TrueUpReviewStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.TrueUpReviewStatus) error); ok {
		r1 = rf(reviewStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTrueUpReviewStatus provides a mock function with given fields: dueDate
func (_m *TrueUpReviewStore) GetTrueUpReviewStatus(dueDate int64) (*model.TrueUpReviewStatus, error) {
	ret := _m.Called(dueDate)

	var r0 *model.TrueUpReviewStatus
	if rf, ok := ret.Get(0).(func(int64) *model.TrueUpReviewStatus); ok {
		r0 = rf(dueDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TrueUpReviewStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(dueDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: reviewStatus
func (_m *TrueUpReviewStore) Update(reviewStatus *model.TrueUpReviewStatus) (*model.TrueUpReviewStatus, error) {
	ret := _m.Called(reviewStatus)

	var r0 *model.TrueUpReviewStatus
	if rf, ok := ret.Get(0).(func(*model.TrueUpReviewStatus) *model.TrueUpReviewStatus); ok {
		r0 = rf(reviewStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TrueUpReviewStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.TrueUpReviewStatus) error); ok {
		r1 = rf(reviewStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
