package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/rodericusifo/echo-template/internal/app/core/employee/controller/request"
	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/pkg/util"
	"github.com/rodericusifo/echo-template/internal/pkg/validator"
	"github.com/rodericusifo/echo-template/pkg/response"
)

func (c *EmployeeController) DeleteEmployee(ctx echo.Context) error {
	reqUser := util.GetRequestUser(ctx)

	reqParams := new(request.DeleteEmployeeRequestParams)
	if err := validator.ValidateRequestParams(ctx, reqParams); err != nil {
		return err
	}

	if err := c.EmployeeService.DeleteEmployee(&input.DeleteEmployeeDTO{
		XID:    reqParams.XID,
		UserID: reqUser.ID,
	}); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.ResponseSuccess[any]("delete employee success", nil, nil))
}
