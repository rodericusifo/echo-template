package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/echo-template/internal/app/core/auth/controller/request"
	"github.com/rodericusifo/echo-template/internal/app/core/auth/controller/response"
	"github.com/rodericusifo/echo-template/internal/app/core/auth/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/app/core/auth/service/dto/output"

	pkg_response "github.com/rodericusifo/echo-template/pkg/response"
)

func init() {
	SetupTestAuthController()
}

func TestAuthController_LoginAuth(t *testing.T) {
	type (
		args struct {
			requestBody request.LoginAuthRequestBody
		}
		result struct {
			responseCode int
			responseBody any
			err          error
		}
	)

	testCases := []struct {
		desc    string
		input   args
		output  result
		before  func()
		after   func()
		isError bool
	}{
		{
			desc: "[ERROR]_because_validation_error",
			input: args{
				requestBody: request.LoginAuthRequestBody{
					Email: mockEmail,
				},
			},
			output:  result{},
			before:  func() {},
			after:   func() {},
			isError: true,
		},
		{
			desc: "[ERROR]_because_unexpected_error_from_service",
			input: args{
				requestBody: request.LoginAuthRequestBody{
					Email:    mockEmail,
					Password: mockPassword,
				},
			},
			output: result{
				err: errors.New("unexpected errors"),
			},
			before: func() {
				{
					var (
						arg1 *input.LoginAuthDTO = &input.LoginAuthDTO{
							Email:    mockEmail,
							Password: mockPassword,
						}
					)
					var (
						result *output.LoginAuthDTO = nil
						err    error                = errors.New("unexpected errors")
					)
					mockAuthService.On("LoginAuth", arg1).Return(result, err).Once()
				}
			},
			after:   func() {},
			isError: true,
		},
		{
			desc: "[SUCCESS]_success_login_auth",
			input: args{
				requestBody: request.LoginAuthRequestBody{
					Email:    mockEmail,
					Password: mockPassword,
				},
			},
			output: result{
				responseCode: http.StatusOK,
				responseBody: pkg_response.ResponseSuccess[any]("auth login success", &response.LoginAuthResponse{
					Token: mockJWTToken,
				}, nil),
			},
			before: func() {
				{
					var (
						arg1 *input.LoginAuthDTO = &input.LoginAuthDTO{
							Email:    mockEmail,
							Password: mockPassword,
						}
					)
					var (
						result *output.LoginAuthDTO = &output.LoginAuthDTO{
							Token: mockJWTToken,
						}
						err error = nil
					)
					mockAuthService.On("LoginAuth", arg1).Return(result, err).Once()
				}
			},
			after:   func() {},
			isError: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			url := "/auth/login"

			strRequestBodyBytes, _ := json.Marshal(tC.input.requestBody)
			strResponseBodyBytes, _ := json.Marshal(tC.output.responseBody)

			req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(string(strRequestBodyBytes)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := mockApp.NewContext(req, rec)

			err := authController.LoginAuth(c)

			if tC.isError {
				assert.Error(t, err)
				if tC.output.err != nil {
					assert.Equal(t, tC.output.err, err)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tC.output.responseCode, rec.Code)
				assert.Equal(t, string(strResponseBodyBytes), strings.TrimSuffix(rec.Body.String(), "\n"))
			}

			tC.after()
		})
	}
}
