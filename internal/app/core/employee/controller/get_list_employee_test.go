package controller

import (
	"encoding/json"
	"errors"
	"fmt"
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
	"github.com/rodericusifo/echo-template/pkg/response/structs"
	"github.com/rodericusifo/echo-template/pkg/util"

	net_url "net/url"

	pkg_response "github.com/rodericusifo/echo-template/pkg/response"
)

func init() {
	SetupTestEmployeeController()
}

func TestEmployeeController_GetListEmployee(t *testing.T) {
	type (
		args struct {
			requestQuery request.GetListEmployeeRequestQuery
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
				requestQuery: request.GetListEmployeeRequestQuery{
					Page: -1,
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
				requestQuery: request.GetListEmployeeRequestQuery{},
			},
			output: result{
				err: errors.New("unexpected errors"),
			},
			before: func() {
				{
					var (
						arg1 *input.GetListEmployeeDTO = &input.GetListEmployeeDTO{
							UserID: 1,
						}
					)
					var (
						result output.GetListEmployeeDTO = nil
						meta   *structs.Meta             = nil
						err    error                     = errors.New("unexpected errors")
					)
					mockEmployeeService.On("GetListEmployee", arg1).Return(result, meta, err).Once()
				}
			},
			after:   func() {},
			isError: true,
		},
		{
			desc: "[SUCCESS]_success_get_list_employee",
			input: args{
				requestQuery: request.GetListEmployeeRequestQuery{},
			},
			output: result{
				responseCode: http.StatusOK,
				responseBody: pkg_response.ResponseSuccess("get list employee success", []*response.EmployeeResponse{
					{
						XID:       mockUUID,
						Name:      "John",
						Email:     "John@gmail.com",
						Address:   &mockAddress,
						Age:       &mockAge,
						Birthday:  &mockBirthday,
						CreatedAt: util.ChangeTypeToTypePointer(mockDate),
						UpdatedAt: util.ChangeTypeToTypePointer(mockDate),
					},
				}, &structs.Meta{
					CurrentPage:      1,
					CountDataPerPage: 1,
					TotalData:        1,
					TotalPage:        1,
				}),
			},
			before: func() {
				{
					var (
						arg1 *input.GetListEmployeeDTO = &input.GetListEmployeeDTO{
							UserID: 1,
						}
					)
					var (
						result output.GetListEmployeeDTO = []*output.EmployeeDTO{
							{
								XID:       mockUUID,
								Name:      "John",
								Email:     "John@gmail.com",
								Address:   &mockAddress,
								Age:       &mockAge,
								Birthday:  &mockBirthday,
								CreatedAt: mockDate,
								UpdatedAt: mockDate,
							},
						}
						meta *structs.Meta = &structs.Meta{
							CurrentPage:      1,
							CountDataPerPage: 1,
							TotalData:        1,
							TotalPage:        1,
						}
						err error = nil
					)
					mockEmployeeService.On("GetListEmployee", arg1).Return(result, meta, err).Once()
				}
			},
			after:   func() {},
			isError: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.before()

			url := "/employees/list"

			q := make(net_url.Values)

			limit := tC.input.requestQuery.Limit
			page := tC.input.requestQuery.Page

			if limit != 0 {
				q.Set("limit", fmt.Sprint(limit))
			}
			if page != 0 {
				q.Set("page", fmt.Sprint(page))
			}

			if len(q) != 0 {
				url = fmt.Sprint(url, "?", q.Encode())
			}

			strResponseBodyBytes, _ := json.Marshal(tC.output.responseBody)

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rec := httptest.NewRecorder()
			c := mockApp.NewContext(req, rec)
			c.Set(constant.C_KEY_REQUEST_USER, mockReqUser)

			err := employeeController.GetListEmployee(c)

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
