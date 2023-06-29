package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/echo-template/pkg/types"
)

func (s *EmployeeService) UpdateEmployee(payload *input.UpdateEmployeeDTO) error {
	employeeModelRes, err := s.EmployeeResource.GetEmployee(&types.Query{
		Searches: [][]types.SearchOperation{
			{
				{Field: "xid", Operator: "=", Value: payload.XID},
				{Field: "user_id", Operator: "=", Value: payload.UserID},
			},
		},
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "employee not found")
		}
		return err
	}

	employeeModel := employeeModelRes

	employeeModel.Address = payload.Address
	employeeModel.Age = payload.Age
	employeeModel.Birthday = payload.Birthday

	err = s.EmployeeResource.UpdateEmployee(employeeModel)
	if err != nil {
		return err
	}

	return nil
}
