// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	input "github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/input"
	mock "github.com/stretchr/testify/mock"

	output "github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/output"

	structs "github.com/rodericusifo/echo-template/pkg/response/structs"
)

// IEmployeeService is an autogenerated mock type for the IEmployeeService type
type IEmployeeService struct {
	mock.Mock
}

// CreateEmployee provides a mock function with given fields: payload
func (_m *IEmployeeService) CreateEmployee(payload *input.CreateEmployeeDTO) error {
	ret := _m.Called(payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(*input.CreateEmployeeDTO) error); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteEmployee provides a mock function with given fields: payload
func (_m *IEmployeeService) DeleteEmployee(payload *input.DeleteEmployeeDTO) error {
	ret := _m.Called(payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(*input.DeleteEmployeeDTO) error); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetEmployee provides a mock function with given fields: payload
func (_m *IEmployeeService) GetEmployee(payload *input.GetEmployeeDTO) (*output.EmployeeDTO, error) {
	ret := _m.Called(payload)

	var r0 *output.EmployeeDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(*input.GetEmployeeDTO) (*output.EmployeeDTO, error)); ok {
		return rf(payload)
	}
	if rf, ok := ret.Get(0).(func(*input.GetEmployeeDTO) *output.EmployeeDTO); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*output.EmployeeDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(*input.GetEmployeeDTO) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListEmployee provides a mock function with given fields: payload
func (_m *IEmployeeService) GetListEmployee(payload *input.GetListEmployeeDTO) ([]*output.EmployeeDTO, *structs.Meta, error) {
	ret := _m.Called(payload)

	var r0 []*output.EmployeeDTO
	var r1 *structs.Meta
	var r2 error
	if rf, ok := ret.Get(0).(func(*input.GetListEmployeeDTO) ([]*output.EmployeeDTO, *structs.Meta, error)); ok {
		return rf(payload)
	}
	if rf, ok := ret.Get(0).(func(*input.GetListEmployeeDTO) []*output.EmployeeDTO); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*output.EmployeeDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(*input.GetListEmployeeDTO) *structs.Meta); ok {
		r1 = rf(payload)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*structs.Meta)
		}
	}

	if rf, ok := ret.Get(2).(func(*input.GetListEmployeeDTO) error); ok {
		r2 = rf(payload)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateEmployee provides a mock function with given fields: payload
func (_m *IEmployeeService) UpdateEmployee(payload *input.UpdateEmployeeDTO) error {
	ret := _m.Called(payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(*input.UpdateEmployeeDTO) error); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIEmployeeService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIEmployeeService creates a new instance of IEmployeeService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIEmployeeService(t mockConstructorTestingTNewIEmployeeService) *IEmployeeService {
	mock := &IEmployeeService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}