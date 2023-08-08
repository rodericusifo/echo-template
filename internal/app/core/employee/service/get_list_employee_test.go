package service

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/output"
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/pkg/response/structs"
	"github.com/rodericusifo/echo-template/pkg/types"
	"github.com/rodericusifo/echo-template/pkg/util"
)

func init() {
	SetupTestEmployeeService()
}

func TestEmployeeService_GetListEmployee(t *testing.T) {
	type (
		args struct {
			payload *input.GetListEmployeeDTO
		}
		result struct {
			value output.GetListEmployeeDTO
			meta  *structs.Meta
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
			desc: "[ERROR]_because_something_error_happens",
			input: args{
				payload: &input.GetListEmployeeDTO{
					Page:   1,
					Limit:  10,
					UserID: 1,
				},
			},
			output: result{
				value: nil,
				meta:  nil,
				err:   errors.New("error something"),
			},
			before: func() {
				{
					var (
						page, limit = util.DefinePageLimitPagination(1, 10)
						offset      = util.CountOffsetPagination(page, limit)

						arg1 *types.Query = &types.Query{
							Limit:  limit,
							Offset: offset,
							Searches: [][]types.SearchOperation{
								{
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result []*sql.Employee = nil
						count  int             = 0
						err    error           = errors.New("error something")
					)
					mockEmployeeResource.On("GetListEmployeeAndCount", arg1).Return(result, count, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_list_employee_not_found",
			input: args{
				payload: &input.GetListEmployeeDTO{
					Page:   1,
					Limit:  10,
					UserID: 1,
				},
			},
			output: result{
				value: nil,
				meta:  nil,
				err:   echo.NewHTTPError(http.StatusNotFound, "list employee not found"),
			},
			before: func() {
				{
					var (
						page, limit = util.DefinePageLimitPagination(1, 10)
						offset      = util.CountOffsetPagination(page, limit)

						arg1 *types.Query = &types.Query{
							Limit:  limit,
							Offset: offset,
							Searches: [][]types.SearchOperation{
								{
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result []*sql.Employee = []*sql.Employee{}
						count  int             = 0
						err    error           = nil
					)
					mockEmployeeResource.On("GetListEmployeeAndCount", arg1).Return(result, count, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_error_on_count_all_employee",
			input: args{
				payload: &input.GetListEmployeeDTO{
					Page:   1,
					Limit:  10,
					UserID: 1,
				},
			},
			output: result{
				value: nil,
				meta:  nil,
				err:   errors.New("error count all employee"),
			},
			before: func() {
				{
					var (
						page, limit = util.DefinePageLimitPagination(1, 10)
						offset      = util.CountOffsetPagination(page, limit)

						arg1 *types.Query = &types.Query{
							Limit:  limit,
							Offset: offset,
							Searches: [][]types.SearchOperation{
								{
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result []*sql.Employee = []*sql.Employee{
							{
								ID:        1,
								XID:       mockUUID,
								Name:      "Someone",
								Email:     "someone@mail.com",
								Address:   nil,
								Age:       nil,
								Birthday:  nil,
								CreatedAt: mockDate,
								UpdatedAt: mockDate,
							},
						}
						count int   = 1
						err   error = nil
					)
					mockEmployeeResource.On("GetListEmployeeAndCount", arg1).Return(result, count, err).Once()
				}
				{
					var (
						arg1 *types.Query = &types.Query{
							Searches: [][]types.SearchOperation{
								{
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						count int   = 0
						err   error = errors.New("error count all employee")
					)
					mockEmployeeResource.On("CountAllEmployee", arg1).Return(count, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_get_list_employee",
			input: args{
				payload: &input.GetListEmployeeDTO{
					Page:   1,
					Limit:  10,
					UserID: 1,
				},
			},
			output: result{
				value: []*output.EmployeeDTO{
					{
						XID:       mockUUID,
						Name:      "Someone",
						Email:     "someone@mail.com",
						Address:   nil,
						Age:       nil,
						Birthday:  nil,
						CreatedAt: mockDate,
						UpdatedAt: mockDate,
					},
				},
				meta: &structs.Meta{
					CurrentPage:      1,
					CountDataPerPage: 1,
					TotalData:        1,
					TotalPage:        1,
				},
				err: nil,
			},
			before: func() {
				{
					var (
						page, limit = util.DefinePageLimitPagination(1, 10)
						offset      = util.CountOffsetPagination(page, limit)

						arg1 *types.Query = &types.Query{
							Limit:  limit,
							Offset: offset,
							Searches: [][]types.SearchOperation{
								{
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result []*sql.Employee = []*sql.Employee{
							{
								ID:        1,
								XID:       mockUUID,
								Name:      "Someone",
								Email:     "someone@mail.com",
								Address:   nil,
								Age:       nil,
								Birthday:  nil,
								CreatedAt: mockDate,
								UpdatedAt: mockDate,
							},
						}
						count int   = 1
						err   error = nil
					)
					mockEmployeeResource.On("GetListEmployeeAndCount", arg1).Return(result, count, err).Once()
				}
				{
					var (
						arg1 *types.Query = &types.Query{
							Searches: [][]types.SearchOperation{
								{
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						count int   = 1
						err   error = nil
					)
					mockEmployeeResource.On("CountAllEmployee", arg1).Return(count, err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			value, meta, err := employeeService.GetListEmployee(tC.input.payload)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.meta, meta)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
