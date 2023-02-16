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

	"github.com/rodericusifo/echo-template/internal/app/core/employee/controller/request"
	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/pkg/response"
)

func init() {
	SetupTestEmployeeController()
}

func TestEmployeeController_CreateEmployee(t *testing.T) {
	type (
		args struct {
			requestBody request.CreateEmployeeRequestBody
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
				requestBody: request.CreateEmployeeRequestBody{
					Name:     "Ifo",
					Email:    "",
					Address:  &mockAddress,
					Age:      &mockAge,
					Birthday: &mockBirthday,
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
				requestBody: request.CreateEmployeeRequestBody{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Address:  &mockAddress,
					Age:      &mockAge,
					Birthday: &mockBirthday,
				},
			},
			output: result{
				err: errors.New("unexpected errors"),
			},
			before: func() {
				{
					var (
						arg1 *input.CreateEmployeeDTO = &input.CreateEmployeeDTO{
							Name:     "Ifo",
							Email:    "Ifo@gmail.com",
							Address:  &mockAddress,
							Age:      &mockAge,
							Birthday: &mockBirthday,
							UserID:   1,
						}
					)
					var (
						err error = errors.New("unexpected errors")
					)
					mockEmployeeService.On("CreateEmployee", arg1).Return(err).Once()
				}
			},
			after:   func() {},
			isError: true,
		},
		{
			desc: "[SUCCESS]_success_create_user",
			input: args{
				requestBody: request.CreateEmployeeRequestBody{
					Name:     "Ifo",
					Email:    "Ifo@gmail.com",
					Address:  &mockAddress,
					Age:      &mockAge,
					Birthday: &mockBirthday,
				},
			},
			output: result{
				responseCode: http.StatusCreated,
				responseBody: response.ResponseSuccess[any]("create employee success", nil, nil),
			},
			before: func() {
				{
					var (
						arg1 *input.CreateEmployeeDTO = &input.CreateEmployeeDTO{
							Name:     "Ifo",
							Email:    "Ifo@gmail.com",
							Address:  &mockAddress,
							Age:      &mockAge,
							Birthday: &mockBirthday,
							UserID:   1,
						}
					)
					var (
						err error = nil
					)
					mockEmployeeService.On("CreateEmployee", arg1).Return(err).Once()
				}
			},
			after:   func() {},
			isError: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			url := "/employees/create"

			strRequestBodyBytes, _ := json.Marshal(tC.input.requestBody)
			strResponseBodyBytes, _ := json.Marshal(tC.output.responseBody)

			req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(string(strRequestBodyBytes)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := mockApp.NewContext(req, rec)
			c.Set(constant.C_KEY_REQUEST_USER, mockReqUser)

			err := employeeController.CreateEmployee(c)

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
