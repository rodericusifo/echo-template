package service

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/output"
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/pkg/types"
)

func init() {
	SetupTestEmployeeService()
}

func TestEmployeeService_GetEmployee(t *testing.T) {
	type (
		args struct {
			payload *input.GetEmployeeDTO
		}
		result struct {
			value *output.EmployeeDTO
			err   error
		}
	)

	testCases := []struct {
		desc   string
		input  args
		output result
		before func()
		after  func()
	}{
		{
			desc: "[ERROR]_because_employee_not_found",
			input: args{
				payload: &input.GetEmployeeDTO{
					XID:    mockUUID,
					UserID: 1,
				},
			},
			output: result{
				value: nil,
				err:   echo.NewHTTPError(http.StatusNotFound, "employee not found"),
			},
			before: func() {
				{
					var (
						arg1 *types.Query = &types.Query{
							Searches: [][]types.SearchOperation{
								{
									{Field: "xid", Operator: "=", Value: mockUUID},
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result *sql.Employee = nil
						err    error         = gorm.ErrRecordNotFound
					)
					mockEmployeeResource.On("GetEmployee", arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_something_error_happens",
			input: args{
				payload: &input.GetEmployeeDTO{
					XID:    mockUUID,
					UserID: 1,
				},
			},
			output: result{
				value: nil,
				err:   errors.New("error something"),
			},
			before: func() {
				{
					var (
						arg1 *types.Query = &types.Query{
							Searches: [][]types.SearchOperation{
								{
									{Field: "xid", Operator: "=", Value: mockUUID},
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result *sql.Employee = nil
						err    error         = errors.New("error something")
					)
					mockEmployeeResource.On("GetEmployee", arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_get_employee",
			input: args{
				payload: &input.GetEmployeeDTO{
					XID:    mockUUID,
					UserID: 1,
				},
			},
			output: result{
				value: &output.EmployeeDTO{
					XID:       mockUUID,
					Name:      "Someone",
					Email:     "someone@mail.com",
					Address:   nil,
					Age:       nil,
					Birthday:  nil,
					CreatedAt: mockDate,
					UpdatedAt: mockDate,
				},
				err: nil,
			},
			before: func() {
				{
					var (
						arg1 *types.Query = &types.Query{
							Searches: [][]types.SearchOperation{
								{
									{Field: "xid", Operator: "=", Value: mockUUID},
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result *sql.Employee = &sql.Employee{
							ID:        1,
							XID:       mockUUID,
							Name:      "Someone",
							Email:     "someone@mail.com",
							Address:   nil,
							Age:       nil,
							Birthday:  nil,
							CreatedAt: mockDate,
							UpdatedAt: mockDate,
						}
						err error = nil
					)
					mockEmployeeResource.On("GetEmployee", arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			value, err := employeeService.GetEmployee(tC.input.payload)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
