// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	sql "github.com/rodericusifo/echo-template/internal/app/model/database/sql"

	types "github.com/rodericusifo/echo-template/pkg/types"
)

// IUserResource is an autogenerated mock type for the IUserResource type
type IUserResource struct {
	mock.Mock
}

// GetUser provides a mock function with given fields: query
func (_m *IUserResource) GetUser(query *types.Query) (*sql.User, error) {
	ret := _m.Called(query)

	var r0 *sql.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*types.Query) (*sql.User, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(*types.Query) *sql.User); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*types.Query) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIUserResource interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUserResource creates a new instance of IUserResource. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUserResource(t mockConstructorTestingTNewIUserResource) *IUserResource {
	mock := &IUserResource{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}