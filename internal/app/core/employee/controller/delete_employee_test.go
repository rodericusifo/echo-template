package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rodericusifo/echo-template/internal/app/core/employee/controller/request"
	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/pkg/response"
)

func init() {
	SetupTestEmployeeController()
}

func TestEmployeeController_DeleteEmployee(t *testing.T) {
	type (
		args struct {
			requestParams request.DeleteEmployeeRequestParams
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
				requestParams: request.DeleteEmployeeRequestParams{},
			},
			output:  result{},
			before:  func() {},
			after:   func() {},
			isError: true,
		},
		{
			desc: "[ERROR]_because_unexpected_error_from_service",
			input: args{
				requestParams: request.DeleteEmployeeRequestParams{
					XID: mockUUID,
				},
			},
			output: result{
				err: errors.New("unexpected errors"),
			},
			before: func() {
				{
					var (
						arg1 *input.DeleteEmployeeDTO = &input.DeleteEmployeeDTO{
							XID:    mockUUID,
							UserID: 1,
						}
					)
					var (
						err error = errors.New("unexpected errors")
					)
					mockEmployeeService.On("DeleteEmployee", arg1).Return(err).Once()
				}
			},
			after:   func() {},
			isError: true,
		},
		{
			desc: "[SUCCESS]_success_delete_employee",
			input: args{
				requestParams: request.DeleteEmployeeRequestParams{
					XID: mockUUID,
				},
			},
			output: result{
				responseCode: http.StatusOK,
				responseBody: response.ResponseSuccess[any]("delete employee success", nil, nil),
			},
			before: func() {
				{
					var (
						arg1 *input.DeleteEmployeeDTO = &input.DeleteEmployeeDTO{
							XID:    mockUUID,
							UserID: 1,
						}
					)
					var (
						err error = nil
					)
					mockEmployeeService.On("DeleteEmployee", arg1).Return(err).Once()
				}
			},
			after:   func() {},
			isError: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			url := "/employees/:id/delete"

			strResponseBodyBytes, _ := json.Marshal(tC.output.responseBody)

			req := httptest.NewRequest(http.MethodDelete, url, nil)
			rec := httptest.NewRecorder()
			c := mockApp.NewContext(req, rec)
			c.Set(constant.C_KEY_REQUEST_USER, mockReqUser)
			c.SetParamNames("xid")
			c.SetParamValues(tC.input.requestParams.XID)

			err := employeeController.DeleteEmployee(c)

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
