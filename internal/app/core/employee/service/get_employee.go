package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/output"
	"github.com/rodericusifo/echo-template/internal/pkg/util/mapper"
	"github.com/rodericusifo/echo-template/pkg/types"
)

func (s *EmployeeService) GetEmployee(payload *input.GetEmployeeDTO) (*output.EmployeeDTO, error) {
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
			return nil, echo.NewHTTPError(http.StatusNotFound, "employee not found")
		}
		return nil, err
	}

	employeeDto := mapper.MapEmployeeToEmployeeDTO(employeeModelRes)

	return employeeDto, nil
}
