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

func (c *EmployeeController) CreateEmployee(ctx echo.Context) error {
	reqUser := util.GetRequestUser(ctx)

	reqBody := new(request.CreateEmployeeRequestBody)
	if err := validator.ValidateRequestBody(ctx, reqBody); err != nil {
		return err
	}

	if err := c.EmployeeService.CreateEmployee(&input.CreateEmployeeDTO{
		Name:     reqBody.Name,
		Email:    reqBody.Email,
		Address:  reqBody.Address,
		Age:      reqBody.Age,
		Birthday: reqBody.Birthday,
		UserID:   reqUser.ID,
	}); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, response.ResponseSuccess[any]("create employee success", nil, nil))
}
