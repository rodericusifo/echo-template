// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IRequestParams is an autogenerated mock type for the IRequestParams type
type IRequestParams struct {
	mock.Mock
}

// CustomValidateRequestParams provides a mock function with given fields:
func (_m *IRequestParams) CustomValidateRequestParams() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIRequestParams interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRequestParams creates a new instance of IRequestParams. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRequestParams(t mockConstructorTestingTNewIRequestParams) *IRequestParams {
	mock := &IRequestParams{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
