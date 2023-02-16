package service

import (
	"errors"
	"net/http"
	"testing"

	"bou.ke/monkey"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/rodericusifo/echo-template/internal/app/core/auth/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/app/core/auth/service/dto/output"
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/internal/pkg/types"
	"github.com/rodericusifo/echo-template/internal/pkg/util"

	pkg_types "github.com/rodericusifo/echo-template/pkg/types"
)

func init() {
	SetupTestAuthService()
}

func TestAuthService_LoginAuth(t *testing.T) {
	type (
		args struct {
			payload *input.LoginAuthDTO
		}
		result struct {
			value *output.LoginAuthDTO
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
			desc: "[ERROR]_because_user_not_found",
			input: args{
				payload: &input.LoginAuthDTO{
					Email:    "someone@mail.com",
					Password: mockPassword,
				},
			},
			output: result{
				value: nil,
				err:   echo.NewHTTPError(http.StatusNotFound, "user not found"),
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.Query = &pkg_types.Query{
							Searches: [][]pkg_types.SearchOperation{
								{
									{Field: "email", Operator: "=", Value: "someone@mail.com"},
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
			desc: "[ERROR]_because_something_error_happens",
			input: args{
				payload: &input.LoginAuthDTO{
					Email:    "someone@mail.com",
					Password: mockPassword,
				},
			},
			output: result{
				value: nil,
				err:   errors.New("error something"),
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.Query = &pkg_types.Query{
							Searches: [][]pkg_types.SearchOperation{
								{
									{Field: "email", Operator: "=", Value: "someone@mail.com"},
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
			desc: "[ERROR]_because_email_and_password_not_match",
			input: args{
				payload: &input.LoginAuthDTO{
					Email:    "someone@mail.com",
					Password: "123",
				},
			},
			output: result{
				value: nil,
				err:   echo.NewHTTPError(http.StatusUnauthorized, "email and password not match"),
			},
			before: func() {
				{
					var (
						arg1 *pkg_types.Query = &pkg_types.Query{
							Searches: [][]pkg_types.SearchOperation{
								{
									{Field: "email", Operator: "=", Value: "someone@mail.com"},
								},
							},
						}
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
					mockUserResource.On("GetUser", arg1).Return(result, err).Once()
				}
			},
			after: func() {},
		},
		{
			desc: "[ERROR]_because_generate_token_from_claims_failed",
			input: args{
				payload: &input.LoginAuthDTO{
					Email:    "someone@mail.com",
					Password: mockPassword,
				},
			},
			output: result{
				value: nil,
				err:   errors.New("error generate token from claims"),
			},
			before: func() {
				{
					monkey.Patch(util.GenerateJWTTokenFromClaims, func(claims *types.JwtCustomClaims) (string, error) {
						return "", errors.New("error generate token from claims")
					})
				}
				{
					var (
						arg1 *pkg_types.Query = &pkg_types.Query{
							Searches: [][]pkg_types.SearchOperation{
								{
									{Field: "email", Operator: "=", Value: "someone@mail.com"},
								},
							},
						}
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
					mockUserResource.On("GetUser", arg1).Return(result, err).Once()
				}
			},
			after: func() {
				{
					monkey.Unpatch(util.GenerateJWTTokenFromClaims)
				}
			},
		},
		{
			desc: "[SUCCESS]_success_login_auth",
			input: args{
				payload: &input.LoginAuthDTO{
					Email:    "someone@mail.com",
					Password: mockPassword,
				},
			},
			output: result{
				value: &output.LoginAuthDTO{
					Token: mockJWTToken,
				},
				err: nil,
			},
			before: func() {
				{
					monkey.Patch(util.GenerateJWTTokenFromClaims, func(claims *types.JwtCustomClaims) (string, error) {
						return mockJWTToken, nil
					})
				}
				{
					var (
						arg1 *pkg_types.Query = &pkg_types.Query{
							Searches: [][]pkg_types.SearchOperation{
								{
									{Field: "email", Operator: "=", Value: "someone@mail.com"},
								},
							},
						}
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
					mockUserResource.On("GetUser", arg1).Return(result, err).Once()
				}
			},
			after: func() {
				{
					monkey.Unpatch(util.GenerateJWTTokenFromClaims)
				}
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			value, err := authService.LoginAuth(tC.input.payload)

			assert.Equal(t, tC.output.err, err)
			assert.Equal(t, tC.output.value, value)

			tC.after()
		})
	}
}
