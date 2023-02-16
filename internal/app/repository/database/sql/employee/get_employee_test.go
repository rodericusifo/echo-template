package employee

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/pkg/types"
)

func init() {
	SetupTestPostgresEmployeeDatabaseSQLRepository()
}

func TestPostgresEmployeeDatabaseSQLRepository_GetEmployee(t *testing.T) {
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
					Selects: []types.SelectOperation{
						{Field: "id"},
						{Field: "xid"},
						{Field: "name"},
						{Field: "email"},
						{Field: "address"},
						{Field: "age"},
						{Field: "birthday"},
						{Field: "created_at"},
						{Field: "updated_at"},
					},
					Searches: [][]types.SearchOperation{
						{
							{Field: "xid", Operator: "=", Value: mockUUID},
							{Field: "email", Operator: "=", Value: "someone@mail.com"},
							{Field: "user_id", Operator: "=", Value: uint(1)},
						},
					},
					Joins: []types.JoinOperation{
						{
							Relation: "User",
						},
					},
					Orders: []types.OrderOperation{
						{Field: "name"},
						{Field: "age", Descending: true},
					},
					Groups: []types.GroupOperation{
						{Field: "name"},
					},
					WithDeleted: true,
				},
			},
			output: result{
				value: nil,
				err:   errors.New("something error"),
			},
			before: func() {
				{
					var (
						arg1 = mockUUID
						arg2 = "someone@mail.com"
						arg3 = 1
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							`SELECT "employees"."id","employees"."xid","employees"."name","employees"."email","employees"."address","employees"."age","employees"."birthday","employees"."created_at","employees"."updated_at","User"."id" AS "User__id","User"."xid" AS "User__xid","User"."name" AS "User__name","User"."email" AS "User__email","User"."password" AS "User__password","User"."role" AS "User__role","User"."created_at" AS "User__created_at","User"."updated_at" AS "User__updated_at" FROM "employees" LEFT JOIN "users" "User" ON "employees"."user_id" = "User"."id" WHERE "employees"."xid" = $1 AND "employees"."email" = $2 AND "employees"."user_id" = $3 GROUP BY "employees"."name" ORDER BY "employees"."name","employees"."age" DESC,"employees"."id" LIMIT 1`,
						),
					).
						WithArgs(arg1, arg2, arg3).
						WillReturnError(errors.New("something error"))
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_something_error_happens_1",
			input: args{
				query: &types.Query{
					Selects: []types.SelectOperation{
						{Field: "id"},
						{Field: "xid"},
						{Field: "name"},
						{Field: "email"},
						{Field: "address"},
						{Field: "age"},
						{Field: "birthday"},
						{Field: "created_at"},
						{Field: "updated_at"},
					},
					Searches: [][]types.SearchOperation{
						{
							{Field: "xid", Operator: "=", Value: mockUUID},
							{Field: "email", Operator: "=", Value: "someone@mail.com"},
							{Field: "user_id", Operator: "=", Value: uint(1)},
						},
					},
					Joins: []types.JoinOperation{
						{
							Relation: "User",
							Selects: []types.SelectOperation{
								{Field: "id"},
								{Field: "xid"},
								{Field: "name"},
							},
						},
					},
				},
			},
			output: result{
				value: nil,
				err:   errors.New("something error"),
			},
			before: func() {
				{
					var (
						arg1 = mockUUID
						arg2 = "someone@mail.com"
						arg3 = 1
					)
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							`SELECT "employees"."id","employees"."xid","employees"."name","employees"."email","employees"."address","employees"."age","employees"."birthday","employees"."created_at","employees"."updated_at","User"."id" AS "User__id","User"."xid" AS "User__xid","User"."name" AS "User__name" FROM "employees" LEFT JOIN "users" "User" ON "employees"."user_id" = "User"."id" WHERE ("employees"."xid" = $1 AND "employees"."email" = $2 AND "employees"."user_id" = $3) AND "employees"."deleted_at" IS NULL ORDER BY "employees"."id" LIMIT 1`,
						),
					).
						WithArgs(arg1, arg2, arg3).
						WillReturnError(errors.New("something error"))
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_get_employee",
			input: args{
				query: &types.Query{
					Distinct: true,
					Selects: []types.SelectOperation{
						{Field: "id"},
						{Field: "xid"},
						{Field: "name"},
						{Field: "email"},
						{Field: "address"},
						{Field: "age"},
						{Field: "birthday"},
						{Field: "created_at"},
						{Field: "updated_at"},
					},
					Searches: [][]types.SearchOperation{
						{
							{Field: "xid", Operator: "=", Value: mockUUID},
							{Field: "email", Operator: "=", Value: "someone@mail.com"},
							{Field: "user_id", Operator: "=", Value: uint(1)},
						},
					},
					Joins: []types.JoinOperation{
						{
							Relation: "User",
							Selects: []types.SelectOperation{
								{Field: "id"},
								{Field: "xid"},
								{Field: "name"},
							},
							Searches: [][]types.SearchOperation{
								{
									{Field: "name", Operator: "LIKE", Value: fmt.Sprint("%", "sometwo", "%")},
								},
							},
						},
					},
				},
			},
			output: result{
				value: &sql.Employee{
					ID:        1,
					XID:       mockUUID,
					Name:      "Someone",
					Email:     "someone@mail.com",
					Address:   nil,
					Age:       nil,
					Birthday:  nil,
					CreatedAt: mockDate,
					UpdatedAt: mockDate,
					User: sql.User{
						ID:   2,
						XID:  mockUUID,
						Name: "sometwo",
					},
				},
				err: nil,
			},
			before: func() {
				{
					var (
						arg1         = fmt.Sprint("%", "sometwo", "%")
						arg2         = mockUUID
						arg3         = "someone@mail.com"
						arg4         = 1
						rowsInstance = sqlmock.NewRows([]string{"id", "xid", "name", "email", "address", "age", "birthday", "created_at", "updated_at", "User__id", "User__xid", "User__name"})
					)
					rowsInstance.AddRow(1, mockUUID, "Someone", "someone@mail.com", nil, nil, nil, mockDate, mockDate, 2, mockUUID, "sometwo")
					mockQuery.ExpectQuery(
						regexp.QuoteMeta(
							`SELECT DISTINCT "employees"."id","employees"."xid","employees"."name","employees"."email","employees"."address","employees"."age","employees"."birthday","employees"."created_at","employees"."updated_at","User"."id" AS "User__id","User"."xid" AS "User__xid","User"."name" AS "User__name" FROM "employees" LEFT JOIN "users" "User" ON "employees"."user_id" = "User"."id" AND "User"."name" LIKE $1 WHERE ("employees"."xid" = $2 AND "employees"."email" = $3 AND "employees"."user_id" = $4) AND "employees"."deleted_at" IS NULL ORDER BY "employees"."id" LIMIT 1`,
						),
					).
						WithArgs(arg1, arg2, arg3, arg4).
						WillReturnRows(rowsInstance)
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			value, err := employeeDatabaseSQLRepository.GetEmployee(tC.input.query)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
