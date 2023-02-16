package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/rodericusifo/echo-template/internal/app/core/employee/controller/request"
	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/pkg/util"
	"github.com/rodericusifo/echo-template/internal/pkg/util/mapper"
	"github.com/rodericusifo/echo-template/internal/pkg/validator"

	pkg_response "github.com/rodericusifo/echo-template/pkg/response"
)

func (c *EmployeeController) GetListEmployee(ctx echo.Context) error {
	reqUser := util.GetRequestUser(ctx)

	reqQuery := new(request.GetListEmployeeRequestQuery)
	if err := validator.ValidateRequestQuery(ctx, reqQuery); err != nil {
		return err
	}

	employeeListDtoRes, meta, err := c.EmployeeService.GetListEmployee(&input.GetListEmployeeDTO{
		Page:   reqQuery.Page,
		Limit:  reqQuery.Limit,
		UserID: reqUser.ID,
	})
	if err != nil {
		return err
	}

	getListEmployeeResponse := mapper.MapEmployeeDTOsToEmployeeResponses(employeeListDtoRes)

	return ctx.JSON(http.StatusOK, pkg_response.ResponseSuccess("get list employee success", getListEmployeeResponse, meta))
}
