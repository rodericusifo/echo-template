package service

import (
	"github.com/rodericusifo/echo-template/internal/app/core/employee/resource"
	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/output"
	"github.com/rodericusifo/echo-template/pkg/response/structs"
)

type IEmployeeService interface {
	CreateEmployee(payload *input.CreateEmployeeDTO) error
	UpdateEmployee(payload *input.UpdateEmployeeDTO) error
	DeleteEmployee(payload *input.DeleteEmployeeDTO) error
	GetListEmployee(payload *input.GetListEmployeeDTO) (output.GetListEmployeeDTO, *structs.Meta, error)
	GetEmployee(payload *input.GetEmployeeDTO) (output.GetEmployeeDTO, error)
}

type EmployeeService struct {
	EmployeeResource resource.IEmployeeResource
}

func InitEmployeeService(employeeResource resource.IEmployeeResource) IEmployeeService {
	return &EmployeeService{
		EmployeeResource: employeeResource,
	}
}
