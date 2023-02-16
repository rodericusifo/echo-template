package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/pkg/types"
)

func (s *EmployeeService) CreateEmployee(payload *input.CreateEmployeeDTO) error {
	employeeModelRes, err := s.EmployeeResource.GetEmployee(&types.Query{
		Selects: []types.SelectOperation{
			{Field: "id"},
		},
		Searches: [][]types.SearchOperation{
			{
				{Field: "email", Operator: "=", Value: payload.Email},
			},
		},
		WithDeleted: true,
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if employeeModelRes != nil {
		return echo.NewHTTPError(http.StatusConflict, "employee already registered")
	}

	employeeModel := &sql.Employee{
		Name:     payload.Name,
		Email:    payload.Email,
		Address:  payload.Address,
		Age:      payload.Age,
		Birthday: payload.Birthday,
		UserID:   payload.UserID,
	}
	err = s.EmployeeResource.CreateEmployee(employeeModel)
	if err != nil {
		return err
	}

	return nil
}
