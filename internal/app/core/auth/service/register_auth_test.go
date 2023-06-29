package service

import (
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/rodericusifo/echo-template/internal/app/core/auth/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/internal/pkg/constant"

	mocks_pkg "github.com/rodericusifo/echo-template/mocks-pkg"
	pkg_types "github.com/rodericusifo/echo-template/pkg/types"
)

func init() {
	SetupTestAuthService()
}

func TestAuthService_RegisterAuth(t *testing.T) {
	type (
		args struct {
			payload *input.RegisterAuthDTO
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
			desc: "[ERROR]_because_error_something_when_get_user",
			input: args{
				payload: &input.RegisterAuthDTO{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Password: "abc123",
					Role:     constant.NON_ADMIN,
				},
			},
			output: result{
				err: errors.New("error something"),
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.Query = &pkg_types.Query{
							Selects: []pkg_types.SelectOperation{
								{Field: "id"},
							},
							Searches: [][]pkg_types.SearchOperation{
								{
									{Field: "email", Operator: "=", Value: "Ifo@gmail.com"},
								},
							},
						}
					)
					var (
						result *sql.User = nil
						err    error     = errors.New("error something")
					)
					mockUserResource.On("GetUser", arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_user_already_registered",
			input: args{
				payload: &input.RegisterAuthDTO{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Password: "abc123",
					Role:     constant.NON_ADMIN,
				},
			},
			output: result{
				err: echo.NewHTTPError(http.StatusConflict, "user already registered"),
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.Query = &pkg_types.Query{
							Selects: []pkg_types.SelectOperation{
								{Field: "id"},
							},
							Searches: [][]pkg_types.SearchOperation{
								{
									{Field: "email", Operator: "=", Value: "Ifo@gmail.com"},
								},
							},
						}
					)
					var (
						result *sql.User = &sql.User{
							ID: 1,
						}
						err error = nil
					)
					mockUserResource.On("GetUser", arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_error_something_when_hash_user_password",
			input: args{
				payload: &input.RegisterAuthDTO{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Password: "abc123",
					Role:     constant.NON_ADMIN,
				},
			},
			output: result{
				err: errors.New("error something"),
			},
			before: func() {
				{
					mocks_pkg.GenerateHashPasswordUtil = func(password string) (string, error) {
						return "", errors.New("error something")
					}
				}
				{
					var (
						arg1 *pkg_types.Query = &pkg_types.Query{
							Selects: []pkg_types.SelectOperation{
								{Field: "id"},
							},
							Searches: [][]pkg_types.SearchOperation{
								{
									{Field: "email", Operator: "=", Value: "Ifo@gmail.com"},
								},
							},
						}
					)
					var (
						result *sql.User = nil
						err    error     = gorm.ErrRecordNotFound
					)
					mockUserResource.On("GetUser", arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_error_something_when_create_user",
			input: args{
				payload: &input.RegisterAuthDTO{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Password: "abc123",
					Role:     constant.NON_ADMIN,
				},
			},
			output: result{
				err: errors.New("error something"),
			},
			before: func() {
				{
					mocks_pkg.GenerateHashPasswordUtil = func(password string) (string, error) {
						return mockHashPassword, nil
					}
				}
				{
					var (
						arg1 *pkg_types.Query = &pkg_types.Query{
							Selects: []pkg_types.SelectOperation{
								{Field: "id"},
							},
							Searches: [][]pkg_types.SearchOperation{
								{
									{Field: "email", Operator: "=", Value: "Ifo@gmail.com"},
								},
							},
						}
					)
					var (
						result *sql.User = nil
						err    error     = gorm.ErrRecordNotFound
					)
					mockUserResource.On("GetUser", arg1).Return(result, err).Once()
				}
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
					mockUserResource.On("CreateUser", arg1).Return(err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[SUCCESS]_success_register_auth",
			input: args{
				payload: &input.RegisterAuthDTO{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Password: "abc123",
					Role:     constant.NON_ADMIN,
				},
			},
			output: result{
				err: nil,
			},
			before: func() {
				{
					mocks_pkg.GenerateHashPasswordUtil = func(password string) (string, error) {
						return mockHashPassword, nil
					}
				}
				{
					var (
						arg1 *pkg_types.Query = &pkg_types.Query{
							Selects: []pkg_types.SelectOperation{
								{Field: "id"},
							},
							Searches: [][]pkg_types.SearchOperation{
								{
									{Field: "email", Operator: "=", Value: "Ifo@gmail.com"},
								},
							},
						}
					)
					var (
						result *sql.User = nil
						err    error     = gorm.ErrRecordNotFound
					)
					mockUserResource.On("GetUser", arg1).Return(result, err).Once()
				}
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
					mockUserResource.On("CreateUser", arg1).Return(err).Once()
				}
			},
			after: func() {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			err := authService.RegisterAuth(tC.input.payload)

			assert.Equal(t, tC.output.err, err)

			tC.after()
		})
	}
}
