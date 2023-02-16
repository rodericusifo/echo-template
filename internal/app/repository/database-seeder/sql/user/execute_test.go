package user

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"

	"github.com/rodericusifo/echo-template/internal/pkg/config"
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/internal/pkg/util"

	mocks_pkg "github.com/rodericusifo/echo-template/mocks-pkg"
)

func SetupTestExecutePostgresUserDatabaseSeederSQLRepository() {
	dialect := constant.POSTGRES
	db, mock := util.MockConnectionDatabaseSQL(dialect)

	mockQuery = mock
	mockDB = db

	layoutFormat := "2006-01-02 15:04:05"
	value := "2015-09-02 08:04:00"
	mockDate, _ = time.Parse(layoutFormat, value)

	mockHashPassword = "$2y$14$rnbG3JhbftD.iQV0QRf5GeNI/XlI85KF2kzrf4hnOs48cSoqPvsmG"
}

func init() {
	SetupTestExecutePostgresUserDatabaseSeederSQLRepository()
}

func TestExecutePostgresUserDatabaseSeederSQLRepository(t *testing.T) {
	type (
		args struct {
			isRebuildData config.IsRebuildDataDBSeederPostgresUser
			db            config.PostgresDBSQLConnection
		}
	)

	testCases := []struct {
		desc   string
		input  args
		before func()
		after  func()
	}{
		{
			desc: "[ERROR]_error_clear_users",
			input: args{
				db:            mockDB,
				isRebuildData: true,
			},
			before: func() {
				{
					var (
						arg1 = constant.ADMIN
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							`SELECT "users"."id" FROM "users" WHERE "users"."role" = $1`,
						),
					).
						WithArgs(arg1).
						WillReturnError(errors.New("error something"))
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_success_clear_users_and_fail_seed_user",
			input: args{
				db:            mockDB,
				isRebuildData: true,
			},
			before: func() {
				{
					var (
						arg1         = constant.ADMIN
						rowsInstance = sqlmock.NewRows([]string{"id"})
					)
					rowsInstance.AddRow(1)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							`SELECT "users"."id" FROM "users" WHERE "users"."role" = $1`,
						),
					).
						WithArgs(arg1).
						WillReturnRows(rowsInstance)
				}
				{
					var (
						arg1 = 1
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							`DELETE FROM "users" WHERE "users"."id" = $1`,
						),
					).
						WithArgs(arg1).
						WillReturnResult(sqlmock.NewResult(0, 1))
					mockQuery.ExpectCommit()
				}
				UserSeedData = []*UserSeedPayload{
					{
						XID:      "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1",
						Name:     "admin",
						Email:    "admin@gmail.com",
						Password: "p4ssw0rd",
						Role:     constant.ADMIN,
					},
				}
				{
					monkey.Patch(time.Now, func() time.Time {
						return mockDate
					})
				}
				{
					mocks_pkg.GenerateHashPasswordUtil = func(password string) (string, error) {
						return mockHashPassword, nil
					}
				}
				{
					var (
						arg1 = "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1"
						arg2 = "admin@gmail.com"
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							`SELECT "users"."id" FROM "users" WHERE "users"."xid" = $1 AND "users"."email" = $2 ORDER BY "users"."id" LIMIT 1`,
						),
					).
						WithArgs(arg1, arg2).
						WillReturnError(gorm.ErrRecordNotFound)
				}
				{
					var (
						arg1 = "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1"
						arg2 = "admin"
						arg3 = "admin@gmail.com"
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
			desc: "[SUCCESS]_success_clear_users_and_seed_users",
			input: args{
				db:            mockDB,
				isRebuildData: true,
			},
			before: func() {
				{
					var (
						arg1         = constant.ADMIN
						rowsInstance = sqlmock.NewRows([]string{"id"})
					)
					rowsInstance.AddRow(1)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							`SELECT "users"."id" FROM "users" WHERE "users"."role" = $1`,
						),
					).
						WithArgs(arg1).
						WillReturnRows(rowsInstance)
				}
				{
					var (
						arg1 = 1
					)
					mockQuery.ExpectBegin()
					mockQuery.ExpectExec(
						regexp.QuoteMeta(
							`DELETE FROM "users" WHERE "users"."id" = $1`,
						),
					).
						WithArgs(arg1).
						WillReturnResult(sqlmock.NewResult(0, 1))
					mockQuery.ExpectCommit()
				}
				UserSeedData = []*UserSeedPayload{
					{
						XID:      "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1",
						Name:     "admin",
						Email:    "admin@gmail.com",
						Password: "p4ssw0rd",
						Role:     constant.ADMIN,
					},
				}
				{
					monkey.Patch(time.Now, func() time.Time {
						return mockDate
					})
				}
				{
					mocks_pkg.GenerateHashPasswordUtil = func(password string) (string, error) {
						return mockHashPassword, nil
					}
				}
				{
					var (
						arg1 = "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1"
						arg2 = "admin@gmail.com"
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							`SELECT "users"."id" FROM "users" WHERE "users"."xid" = $1 AND "users"."email" = $2 ORDER BY "users"."id" LIMIT 1`,
						),
					).
						WithArgs(arg1, arg2).
						WillReturnError(gorm.ErrRecordNotFound)
				}
				{
					var (
						arg1         = "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1"
						arg2         = "admin"
						arg3         = "admin@gmail.com"
						arg4         = mockHashPassword
						arg5         = constant.ADMIN
						arg6         = mockDate.Local()
						arg7         = mockDate.Local()
						rowsInstance = sqlmock.NewRows([]string{"id"})
					)
					rowsInstance.AddRow(1)
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
					monkey.Unpatch(time.Now)
				}
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			ExecutePostgresUserDatabaseSeederRepository(tC.input.isRebuildData, tC.input.db)

			tC.after()
		})
	}
}
