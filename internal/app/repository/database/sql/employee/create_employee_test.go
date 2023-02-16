package employee

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
)

func init() {
	SetupTestPostgresEmployeeDatabaseSQLRepository()
}

func TestPostgresEmployeeDatabaseSQLRepository_CreateEmployee(t *testing.T) {
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
			desc: "[ERROR]_because_missing_email",
			input: args{
				payload: &sql.Employee{
					Name:   "Someone",
					UserID: 1,
				},
			},
			output: result{
				err: errors.New("missing email employee"),
			},
			before: func() {
				{
					monkey.Patch(uuid.NewString, func() string {
						return mockUUID
					})
				}
				{
					monkey.Patch(time.Now, func() time.Time {
						return mockDate
					})
				}
				{
					var (
						arg1      = mockUUID
						arg2      = "Someone"
						arg3      = ""
						arg4  any = nil
						arg5  any = nil
						arg6  any = nil
						arg7      = 1
						arg8      = mockDate.Local()
						arg9      = mockDate.Local()
						arg10 any = nil
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							`INSERT INTO "employees" ("xid","name","email","address","age","birthday","user_id","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING "id"`,
						),
					).
						WithArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10).
						WillReturnError(errors.New("missing email employee"))
					mockQuery.ExpectRollback()
				}
			},
			after: func() {
				{
					monkey.Unpatch(uuid.NewString)
				}
				{
					monkey.Unpatch(time.Now)
				}
			},
		},
		{
			desc: "[SUCCESS]_success_create_employee",
			input: args{
				payload: &sql.Employee{
					Name:   "Someone",
					Email:  "someone@mail.com",
					UserID: 1,
				},
			},
			output: result{
				err: nil,
			},
			before: func() {
				{
					monkey.Patch(uuid.NewString, func() string {
						return mockUUID
					})
				}
				{
					monkey.Patch(time.Now, func() time.Time {
						return mockDate
					})
				}
				{
					var (
						arg1             = mockUUID
						arg2             = "Someone"
						arg3             = "someone@mail.com"
						arg4         any = nil
						arg5         any = nil
						arg6         any = nil
						arg7             = 1
						arg8             = mockDate.Local()
						arg9             = mockDate.Local()
						arg10        any = nil
						rowsInstance     = sqlmock.NewRows([]string{"id"})
					)
					rowsInstance.AddRow(1)
					mockQuery.ExpectBegin()
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							`INSERT INTO "employees" ("xid","name","email","address","age","birthday","user_id","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING "id"`,
						),
					).
						WithArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10).
						WillReturnRows(rowsInstance)
					mockQuery.ExpectCommit()
				}
			},
			after: func() {
				{
					monkey.Unpatch(uuid.NewString)
				}
				{
					monkey.Unpatch(time.Now)
				}
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			err := employeeDatabaseSQLRepository.CreateEmployee(tC.input.payload)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
