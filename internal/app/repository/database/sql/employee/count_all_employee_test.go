package employee

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/echo-template/pkg/types"
)

func init() {
	SetupTestPostgresEmployeeDatabaseSQLRepository()
}

func TestPostgresEmployeeDatabaseSQLRepository_CountAllEmployee(t *testing.T) {
	type (
		args struct {
			query *types.Query
		}
		result struct {
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
					Searches: [][]types.SearchOperation{
						{
							{Field: "user_id", Operator: "=", Value: uint(1)},
						},
					},
				},
			},
			output: result{
				count: 0,
				err:   errors.New("something error"),
			},
			before: func() {
				{
					var (
						arg1 = 1
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							`SELECT "employees"."id" FROM "employees" WHERE "employees"."user_id" = $1 AND "employees"."deleted_at" IS NULL`,
						),
					).
						WithArgs(arg1).
						WillReturnError(errors.New("something error"))
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
				count: 1,
				err:   nil,
			},
			before: func() {
				{
					var (
						arg1         = 1
						rowsInstance = sqlmock.NewRows([]string{"id"})
					)
					rowsInstance.AddRow(1)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							`SELECT "employees"."id" FROM "employees" WHERE "employees"."user_id" = $1 AND "employees"."deleted_at" IS NULL`,
						),
					).
						WithArgs(arg1).
						WillReturnRows(rowsInstance)
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			count, err := employeeDatabaseSQLRepository.CountAllEmployee(tC.input.query)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.count, count)

			tC.after()
		})
	}
}
