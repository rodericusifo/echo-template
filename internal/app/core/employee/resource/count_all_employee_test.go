package resource

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/echo-template/pkg/types"
)

func init() {
	SetupTestEmployeeResource()
}

func TestEmployeeResource_CountAllEmployee(t *testing.T) {
	type (
		args struct {
			query *types.Query
		}
		result struct {
			value int
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
					Searches: [][]types.SearchOperation{
						{
							{Field: "user_id", Operator: "=", Value: uint(1)},
						},
					},
				},
			},
			output: result{
				value: 0,
				err:   errors.New("error something"),
			},
			before: func() {
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
						result int   = 0
						err    error = errors.New("error something")
					)
					mockEmployeeDatabaseSQLRepository.On("CountAllEmployee", arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_count_all_employee",
			input: args{
				query: &types.Query{
					Searches: [][]types.SearchOperation{
						{
							{Field: "user_id", Operator: "=", Value: uint(1)},
						},
					},
				},
			},
			output: result{
				value: 1,
				err:   nil,
			},
			before: func() {
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
						result int   = 1
						err    error = nil
					)
					mockEmployeeDatabaseSQLRepository.On("CountAllEmployee", arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			value, err := employeeResource.CountAllEmployee(tC.input.query)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
