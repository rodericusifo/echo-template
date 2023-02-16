package user

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
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
)

func init() {
	SetupTestPostgresUserDatabaseSQLRepository()
}

func TestPostgresUserDatabaseSQLRepository_CreateUser(t *testing.T) {
	type (
		args struct {
			payload *sql.User
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
				payload: &sql.User{
					Name:     "Someone",
					Password: mockHashPassword,
					Role:     constant.ADMIN,
				},
			},
			output: result{
				err: errors.New("missing email user"),
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
						arg1 = mockUUID
						arg2 = "Someone"
						arg3 = ""
						arg4 = mockHashPassword
						arg5 = constant.ADMIN
						arg6 = mockDate.Local()
						arg7 = mockDate.Local()
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							`INSERT INTO "users" ("xid","name","email","password","role","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`,
						),
					).
						WithArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7).
						WillReturnError(errors.New("missing email user"))
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
			desc: "[SUCCESS]_success_create_user",
			input: args{
				payload: &sql.User{
					Name:     "Someone",
					Email:    "someone@mail.com",
					Password: mockHashPassword,
					Role:     constant.ADMIN,
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
						arg1         = mockUUID
						arg2         = "Someone"
						arg3         = "someone@mail.com"
						arg4         = mockHashPassword
						arg5         = constant.ADMIN
						arg6         = mockDate.Local()
						arg7         = mockDate.Local()
						rowsInstance = sqlmock.NewRows([]string{"id"})
					)
					rowsInstance.AddRow(3)
					mockQuery.ExpectBegin()
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							`INSERT INTO "users" ("xid","name","email","password","role","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`,
						),
					).
						WithArgs(arg1, arg2, arg3, arg4, arg5, arg6, arg7).
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

			err := userDatabaseSQLRepository.CreateUser(tC.input.payload)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
