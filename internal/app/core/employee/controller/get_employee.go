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

func (c *EmployeeController) GetEmployee(ctx echo.Context) error {
	reqUser := util.GetRequestUser(ctx)

	reqParams := new(request.GetEmployeeRequestParams)
	if err := validator.ValidateRequestParams(ctx, reqParams); err != nil {
		return err
	}

	employeeDtoRes, err := c.EmployeeService.GetEmployee(&input.GetEmployeeDTO{
		XID:    reqParams.XID,
		UserID: reqUser.ID,
	})
	if err != nil {
		return err
	}

	getEmployeeResponse := mapper.MapEmployeeDTOToEmployeeResponse(employeeDtoRes)

	return ctx.JSON(http.StatusOK, pkg_response.ResponseSuccess("get employee success", getEmployeeResponse, nil))
}
