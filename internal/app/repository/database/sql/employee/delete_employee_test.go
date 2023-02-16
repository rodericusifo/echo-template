package employee

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
)

func init() {
	SetupTestPostgresEmployeeDatabaseSQLRepository()
}

func TestPostgresEmployeeDatabaseSQLRepository_DeleteEmployee(t *testing.T) {
	type (
		args struct {
			payload *sql.Employee
		}
		result struct {
			err error
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
				payload: &sql.Employee{
					ID: 1,
				},
			},
			output: result{
				err: errors.New("error something"),
			},
			before: func() {
				{
					monkey.Patch(time.Now, func() time.Time {
						return mockDate
					})
				}
				{
					var (
						arg1 = mockDate.Local()
						arg2 = 1
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							`UPDATE "employees" SET "deleted_at"=$1 WHERE "employees"."id" = $2 AND "employees"."deleted_at" IS NULL`,
						),
					).
						WithArgs(arg1, arg2).
						WillReturnError(errors.New("error something"))
					mockQuery.ExpectRollback()
				}
			},
			after: func() {
				{
					monkey.Unpatch(time.Now)
				}
			},
		},
		{
			desc: "[SUCCESS]_success_delete_employee",
			input: args{
				payload: &sql.Employee{
					ID: 1,
				},
			},
			output: result{
				err: nil,
			},
			before: func() {
				{
					monkey.Patch(time.Now, func() time.Time {
						return mockDate
					})
				}
				{
					var (
						arg1 = mockDate.Local()
						arg2 = 1
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							`UPDATE "employees" SET "deleted_at"=$1 WHERE "employees"."id" = $2 AND "employees"."deleted_at" IS NULL`,
						),
					).
						WithArgs(arg1, arg2).
						WillReturnResult(sqlmock.NewResult(0, 1))
					mockQuery.ExpectCommit()
				}
			},
			after: func() {
				{
					monkey.Unpatch(time.Now)
				}
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			err := employeeDatabaseSQLRepository.DeleteEmployee(tC.input.payload)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
