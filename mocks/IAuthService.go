// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	input "github.com/rodericusifo/echo-template/internal/app/core/auth/service/dto/input"
	mock "github.com/stretchr/testify/mock"

	output "github.com/rodericusifo/echo-template/internal/app/core/auth/service/dto/output"
)

// IAuthService is an autogenerated mock type for the IAuthService type
type IAuthService struct {
	mock.Mock
}

// LoginAuth provides a mock function with given fields: payload
func (_m *IAuthService) LoginAuth(payload *input.LoginAuthDTO) (*output.LoginAuthDTO, error) {
	ret := _m.Called(payload)

	var r0 *output.LoginAuthDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(*input.LoginAuthDTO) (*output.LoginAuthDTO, error)); ok {
		return rf(payload)
	}
	if rf, ok := ret.Get(0).(func(*input.LoginAuthDTO) *output.LoginAuthDTO); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*output.LoginAuthDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(*input.LoginAuthDTO) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterAuth provides a mock function with given fields: payload
func (_m *IAuthService) RegisterAuth(payload *input.RegisterAuthDTO) error {
	ret := _m.Called(payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(*input.RegisterAuthDTO) error); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIAuthService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIAuthService creates a new instance of IAuthService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIAuthService(t mockConstructorTestingTNewIAuthService) *IAuthService {
	mock := &IAuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
