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

func (c *EmployeeController) UpdateEmployee(ctx echo.Context) error {
	reqUser := util.GetRequestUser(ctx)

	reqBody := new(request.UpdateEmployeeRequestBody)
	if err := validator.ValidateRequestBody(ctx, reqBody); err != nil {
		return err
	}

	reqParams := new(request.UpdateEmployeeRequestParams)
	if err := validator.ValidateRequestParams(ctx, reqParams); err != nil {
		return err
	}

	if err := c.EmployeeService.UpdateEmployee(&input.UpdateEmployeeDTO{
		XID:      reqParams.XID,
		Address:  reqBody.Address,
		Age:      reqBody.Age,
		Birthday: reqBody.Birthday,
		UserID:   reqUser.ID,
	}); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.ResponseSuccess[any]("update employee success", nil, nil))
}
