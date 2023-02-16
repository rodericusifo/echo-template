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

func TestEmployeeController_UpdateEmployee(t *testing.T) {
	type (
		args struct {
			requestParams request.UpdateEmployeeRequestParams
			requestBody   request.UpdateEmployeeRequestBody
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
			desc: "[ERROR]_because_validation_error_body",
			input: args{
				requestParams: request.UpdateEmployeeRequestParams{
					XID: mockUUID,
				},
				requestBody: request.UpdateEmployeeRequestBody{
					Address:  &mockAddress,
					Age:      &mockAgeMinus,
					Birthday: &mockBirthday,
				},
			},
			output:  result{},
			before:  func() {},
			after:   func() {},
			isError: true,
		},
		{
			desc: "[ERROR]_because_validation_error_params",
			input: args{
				requestParams: request.UpdateEmployeeRequestParams{},
				requestBody: request.UpdateEmployeeRequestBody{
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
				requestParams: request.UpdateEmployeeRequestParams{
					XID: mockUUID,
				},
				requestBody: request.UpdateEmployeeRequestBody{
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
						arg1 *input.UpdateEmployeeDTO = &input.UpdateEmployeeDTO{
							XID:      mockUUID,
							Address:  &mockAddress,
							Age:      &mockAge,
							Birthday: &mockBirthday,
							UserID:   1,
						}
					)
					var (
						err error = errors.New("unexpected errors")
					)
					mockEmployeeService.On("UpdateEmployee", arg1).Return(err).Once()
				}
			},
			after:   func() {},
			isError: true,
		},
		{
			desc: "[SUCCESS]_success_update_employee",
			input: args{
				requestParams: request.UpdateEmployeeRequestParams{
					XID: mockUUID,
				},
				requestBody: request.UpdateEmployeeRequestBody{
					Address:  &mockAddress,
					Age:      &mockAge,
					Birthday: &mockBirthday,
				},
			},
			output: result{
				responseCode: http.StatusOK,
				responseBody: response.ResponseSuccess[any]("update employee success", nil, nil),
			},
			before: func() {
				{
					var (
						arg1 *input.UpdateEmployeeDTO = &input.UpdateEmployeeDTO{
							XID:      mockUUID,
							Address:  &mockAddress,
							Age:      &mockAge,
							Birthday: &mockBirthday,
							UserID:   1,
						}
					)
					var (
						err error = nil
					)
					mockEmployeeService.On("UpdateEmployee", arg1).Return(err).Once()
				}
			},
			after:   func() {},
			isError: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			url := "/employees/:id/update"

			strRequestBodyBytes, _ := json.Marshal(tC.input.requestBody)
			strResponseBodyBytes, _ := json.Marshal(tC.output.responseBody)

			req := httptest.NewRequest(http.MethodPatch, url, strings.NewReader(string(strRequestBodyBytes)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := mockApp.NewContext(req, rec)
			c.Set(constant.C_KEY_REQUEST_USER, mockReqUser)
			c.SetParamNames("xid")
			c.SetParamValues(tC.input.requestParams.XID)

			err := employeeController.UpdateEmployee(c)

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
