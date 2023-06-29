package resource

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
)

func init() {
	SetupTestUserResource()
}

func TestEmployeeResource_CreateEmployee(t *testing.T) {
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
			desc: "[ERROR]_because_something_error_happens",
			input: args{
				payload: &sql.User{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Password: mockHashPassword,
					Role:     constant.NON_ADMIN,
				},
			},
			output: result{
				err: errors.New("error something"),
			},
			before: func() {
				{
					var (
						arg1 *sql.User = &sql.User{
							Name:     "Ifo",
							Email:    "Ifo@gmail.com",
							Password: mockHashPassword,
							Role:     constant.NON_ADMIN,
						}
					)
					var (
						err error = errors.New("error something")
					)
					mockUserDatabaseSQLRepository.On("CreateUser", arg1).Return(err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_create_employee",
			input: args{
				payload: &sql.User{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Password: mockHashPassword,
					Role:     constant.NON_ADMIN,
				},
			},
			output: result{
				err: nil,
			},
			before: func() {
				{
					var (
						arg1 *sql.User = &sql.User{
							Name:     "Ifo",
							Email:    "Ifo@gmail.com",
							Password: mockHashPassword,
							Role:     constant.NON_ADMIN,
						}
					)
					var (
						err error = nil
					)
					mockUserDatabaseSQLRepository.On("CreateUser", arg1).Return(err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			err := userResource.CreateUser(tC.input.payload)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
