package resource

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/pkg/types"
)

func init() {
	SetupTestEmployeeResource()
}

func TestEmployeeResource_GetEmployee(t *testing.T) {
	type (
		args struct {
			query *types.Query
		}
		result struct {
			value *sql.Employee
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
							{Field: "xid", Operator: "=", Value: mockUUID},
							{Field: "email", Operator: "=", Value: "someone@mail.com"},
							{Field: "user_id", Operator: "=", Value: uint(1)},
						},
					},
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
									{Field: "email", Operator: "=", Value: "someone@mail.com"},
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result *sql.Employee = nil
						err    error         = errors.New("error something")
					)
					mockEmployeeDatabaseSQLRepository.On("GetEmployee", arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_get_employee",
			input: args{
				query: &types.Query{
					Searches: [][]types.SearchOperation{
						{
							{Field: "xid", Operator: "=", Value: mockUUID},
							{Field: "email", Operator: "=", Value: "someone@mail.com"},
							{Field: "user_id", Operator: "=", Value: uint(1)},
						},
					},
				},
			},
			output: result{
				value: &sql.Employee{
					ID:        1,
					XID:       mockUUID,
					Name:      "Ifo",
					Email:     "Ifo@gmail.com",
					Address:   &mockAddress,
					Age:       &mockAge,
					Birthday:  &mockBirthday,
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
									{Field: "email", Operator: "=", Value: "someone@mail.com"},
									{Field: "user_id", Operator: "=", Value: uint(1)},
								},
							},
						}
					)
					var (
						result *sql.Employee = &sql.Employee{
							ID:        1,
							XID:       mockUUID,
							Name:      "Ifo",
							Email:     "Ifo@gmail.com",
							Address:   &mockAddress,
							Age:       &mockAge,
							Birthday:  &mockBirthday,
							CreatedAt: mockDate,
							UpdatedAt: mockDate,
						}
						err error = nil
					)
					mockEmployeeDatabaseSQLRepository.On("GetEmployee", arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			value, err := employeeResource.GetEmployee(tC.input.query)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
