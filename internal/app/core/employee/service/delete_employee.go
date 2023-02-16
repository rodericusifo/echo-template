package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/echo-template/pkg/types"
)

func (s *EmployeeService) DeleteEmployee(payload *input.DeleteEmployeeDTO) error {
	employeeModelRes, err := s.EmployeeResource.GetEmployee(&types.Query{
		Selects: []types.SelectOperation{
			{Field: "id"},
		},
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

	err = s.EmployeeResource.DeleteEmployee(employeeModel)
	if err != nil {
		return err
	}

	return nil
}
