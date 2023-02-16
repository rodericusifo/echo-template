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

func TestPostgresEmployeeDatabaseSQLRepository_UpdateEmployee(t *testing.T) {
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
					ID:        1,
					XID:       mockUUID,
					Name:      "Someone",
					Email:     "someone@mail.com",
					Address:   &mockAddress,
					Age:       &mockAge,
					Birthday:  &mockBirthday,
					UserID:    1,
					CreatedAt: mockDate,
					UpdatedAt: mockDate.Local(),
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
						arg1      = mockUUID
						arg2      = "Someone"
						arg3      = "someone@mail.com"
						arg4      = &mockAddress
						arg5      = &mockAge
						arg6      = &mockBirthday
						arg7      = 1
						arg8      = mockDate
						arg9      = mockDate.Local()
						arg10 any = nil
						arg11     = 1
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							`UPDATE "employees" SET "xid"=$1,"name"=$2,"email"=$3,"address"=$4,"age"=$5,"birthday"=$6,"user_id"=$7,"created_at"=$8,"updated_at"=$9,"deleted_at"=$10 WHERE "employees"."deleted_at" IS NULL AND "id" = $11`,
						),
					).
						WithArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11).
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
			desc: "[SUCCESS]_success_update_employee",
			input: args{
				payload: &sql.Employee{
					ID:        1,
					XID:       mockUUID,
					Name:      "Someone",
					Email:     "someone@mail.com",
					Address:   &mockAddress,
					Age:       &mockAge,
					Birthday:  &mockBirthday,
					UserID:    1,
					CreatedAt: mockDate,
					UpdatedAt: mockDate.Local(),
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
						arg1      = mockUUID
						arg2      = "Someone"
						arg3      = "someone@mail.com"
						arg4      = &mockAddress
						arg5      = &mockAge
						arg6      = &mockBirthday
						arg7      = 1
						arg8      = mockDate
						arg9      = mockDate.Local()
						arg10 any = nil
						arg11     = 1
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							`UPDATE "employees" SET "xid"=$1,"name"=$2,"email"=$3,"address"=$4,"age"=$5,"birthday"=$6,"user_id"=$7,"created_at"=$8,"updated_at"=$9,"deleted_at"=$10 WHERE "employees"."deleted_at" IS NULL AND "id" = $11`,
						),
					).
						WithArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11).
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

			err := employeeDatabaseSQLRepository.UpdateEmployee(tC.input.payload)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
