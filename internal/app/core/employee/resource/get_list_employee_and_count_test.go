package resource

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/pkg/types"
	"github.com/rodericusifo/echo-template/pkg/util"
)

func init() {
	SetupTestEmployeeResource()
}

func TestEmployeeResource_GetEmployeeListAndCount(t *testing.T) {
	type (
		args struct {
			query *types.Query
		}
		result struct {
			value []*sql.Employee
			count int
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
				query: &types.Query{
					Offset: util.CountOffsetPagination(1, 10),
					Limit:  10,
					Searches: [][]types.SearchOperation{
						{
							{Field: "user_id", Operator: "=", Value: uint(1)},
						},
					},
				},
			},
			output: result{
				value: nil,
				count: 0,
				err:   errors.New("error something"),
			},
			before: func() {
				{
					var (
						arg1 *types.Query = &types.Query{
							Offset: util.CountOffsetPagination(1, 10),
							Limit:  10,
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
					mockEmployeeDatabaseSQLRepository.On("GetListEmployeeAndCount", arg1).Return(result, count, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_get_list_employee_and_count",
			input: args{
				query: &types.Query{
					Offset: util.CountOffsetPagination(1, 10),
					Limit:  10,
					Searches: [][]types.SearchOperation{
						{
							{Field: "user_id", Operator: "=", Value: uint(1)},
						},
					},
				},
			},
			output: result{
				value: []*sql.Employee{
					{
						ID:        1,
						XID:       mockUUID,
						Name:      "Ifo",
						Email:     "Ifo@gmail.com",
						Address:   &mockAddress,
						Age:       &mockAge,
						Birthday:  &mockBirthday,
						UserID:    1,
						CreatedAt: mockDate,
						UpdatedAt: mockDate,
					},
				},
				count: 1,
				err:   nil,
			},
			before: func() {
				{
					var (
						arg1 *types.Query = &types.Query{
							Offset: util.CountOffsetPagination(1, 10),
							Limit:  10,
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
								Name:      "Ifo",
								Email:     "Ifo@gmail.com",
								Address:   &mockAddress,
								Age:       &mockAge,
								Birthday:  &mockBirthday,
								UserID:    1,
								CreatedAt: mockDate,
								UpdatedAt: mockDate,
							},
						}
						count int   = 1
						err   error = nil
					)
					mockEmployeeDatabaseSQLRepository.On("GetListEmployeeAndCount", arg1).Return(result, count, err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			value, count, err := employeeResource.GetListEmployeeAndCount(tC.input.query)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.count, count)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
