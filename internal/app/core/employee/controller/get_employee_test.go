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
	"github.com/rodericusifo/echo-template/internal/app/core/employee/controller/response"
	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/output"
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/pkg/util"

	pkg_response "github.com/rodericusifo/echo-template/pkg/response"
)

func init() {
	SetupTestEmployeeController()
}

func TestEmployeeController_GetEmployee(t *testing.T) {
	type (
		args struct {
			requestParams request.GetEmployeeRequestParams
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
				requestParams: request.GetEmployeeRequestParams{},
			},
			output:  result{},
			before:  func() {},
			after:   func() {},
			isError: true,
		},
		{
			desc: "[ERROR]_because_unexpected_error_from_service",
			input: args{
				requestParams: request.GetEmployeeRequestParams{
					XID: mockUUID,
				},
			},
			output: result{
				err: errors.New("unexpected errors"),
			},
			before: func() {
				{
					var (
						arg1 *input.GetEmployeeDTO = &input.GetEmployeeDTO{
							XID:    mockUUID,
							UserID: 1,
						}
					)
					var (
						result output.GetEmployeeDTO = nil
						err    error                 = errors.New("unexpected errors")
					)
					mockEmployeeService.On("GetEmployee", arg1).Return(result, err).Once()
				}
			},
			after:   func() {},
			isError: true,
		},
		{
			desc: "[SUCCESS]_success_get_employee",
			input: args{
				requestParams: request.GetEmployeeRequestParams{
					XID: mockUUID,
				},
			},
			output: result{
				responseCode: http.StatusOK,
				responseBody: pkg_response.ResponseSuccess("get employee success", &response.EmployeeResponse{
					XID:       mockUUID,
					Name:      "John",
					Email:     "John@gmail.com",
					Address:   &mockAddress,
					Age:       &mockAge,
					Birthday:  &mockBirthday,
					CreatedAt: util.ChangeTypeToTypePointer(mockDate),
					UpdatedAt: util.ChangeTypeToTypePointer(mockDate),
				}, nil),
			},
			before: func() {
				{
					var (
						arg1 *input.GetEmployeeDTO = &input.GetEmployeeDTO{
							XID:    mockUUID,
							UserID: 1,
						}
					)
					var (
						result output.GetEmployeeDTO = &output.EmployeeDTO{
							XID:       mockUUID,
							Name:      "John",
							Email:     "John@gmail.com",
							Address:   &mockAddress,
							Age:       &mockAge,
							Birthday:  &mockBirthday,
							CreatedAt: mockDate,
							UpdatedAt: mockDate,
						}
						err error = nil
					)
					mockEmployeeService.On("GetEmployee", arg1).Return(result, err).Once()
				}
			},
			after:   func() {},
			isError: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			url := "/employees/:id/detail"

			strResponseBodyBytes, _ := json.Marshal(tC.output.responseBody)

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rec := httptest.NewRecorder()
			c := mockApp.NewContext(req, rec)
			c.Set(constant.C_KEY_REQUEST_USER, mockReqUser)
			c.SetParamNames("xid")
			c.SetParamValues(tC.input.requestParams.XID)

			err := employeeController.GetEmployee(c)

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
