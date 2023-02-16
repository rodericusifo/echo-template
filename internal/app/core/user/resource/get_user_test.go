package resource

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/pkg/types"
)

func init() {
	SetupTestUserResource()
}

func TestUserResource_GetUser(t *testing.T) {
	type (
		args struct {
			query *types.Query
		}
		result struct {
			value *sql.User
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
				query: nil,
			},
			output: result{
				value: nil,
				err:   errors.New("error something"),
			},
			before: func() {
				{
					var (
						arg1 *types.Query = nil
					)
					var (
						result *sql.User = nil
						err    error     = errors.New("error something")
					)
					mockUserDatabaseSQLRepository.On("GetUser", arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_get_user",
			input: args{
				query: nil,
			},
			output: result{
				value: &sql.User{
					ID:        3,
					XID:       mockUUID,
					Name:      "Someone",
					Email:     "someone@mail.com",
					Password:  mockHashPassword,
					Role:      constant.ADMIN,
					CreatedAt: mockDate,
					UpdatedAt: mockDate,
				},
				err: nil,
			},
			before: func() {
				{
					var (
						arg1 *types.Query = nil
					)
					var (
						result *sql.User = &sql.User{
							ID:        3,
							XID:       mockUUID,
							Name:      "Someone",
							Email:     "someone@mail.com",
							Password:  mockHashPassword,
							Role:      constant.ADMIN,
							CreatedAt: mockDate,
							UpdatedAt: mockDate,
						}
						err error = nil
					)
					mockUserDatabaseSQLRepository.On("GetUser", arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			value, err := userResource.GetUser(tC.input.query)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
