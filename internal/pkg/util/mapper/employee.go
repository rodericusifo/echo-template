package mapper

import (
	"github.com/rodericusifo/echo-template/internal/app/core/employee/controller/response"
	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/output"
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/pkg/util"
)

func MapEmployeeToEmployeeDTO(model *sql.Employee) *output.EmployeeDTO {
	return &output.EmployeeDTO{
		XID:       model.XID,
		Name:      model.Name,
		Email:     model.Email,
		Address:   model.Address,
		Age:       model.Age,
		Birthday:  model.Birthday,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func MapEmployeesToEmployeeDTOs(models []*sql.Employee) []*output.EmployeeDTO {
	result := make([]*output.EmployeeDTO, 0)

	for _, model := range models {
		result = append(result, &output.EmployeeDTO{
			XID:       model.XID,
			Name:      model.Name,
			Email:     model.Email,
			Address:   model.Address,
			Age:       model.Age,
			Birthday:  model.Birthday,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		})
	}

	return result
}

func MapEmployeeDTOToEmployeeResponse(dto *output.EmployeeDTO) *response.EmployeeResponse {
	return &response.EmployeeResponse{
		XID:       dto.XID,
		Name:      dto.Name,
		Email:     dto.Email,
		Address:   dto.Address,
		Age:       dto.Age,
		Birthday:  dto.Birthday,
		CreatedAt: util.ChangeTypeToTypePointer(dto.CreatedAt),
		UpdatedAt: util.ChangeTypeToTypePointer(dto.UpdatedAt),
	}
}

func MapEmployeeDTOsToEmployeeResponses(dtos []*output.EmployeeDTO) []*response.EmployeeResponse {
	result := make([]*response.EmployeeResponse, 0)

	for _, dto := range dtos {
		result = append(result, &response.EmployeeResponse{
			XID:       dto.XID,
			Name:      dto.Name,
			Email:     dto.Email,
			Address:   dto.Address,
			Age:       dto.Age,
			Birthday:  dto.Birthday,
			CreatedAt: util.ChangeTypeToTypePointer(dto.CreatedAt),
			UpdatedAt: util.ChangeTypeToTypePointer(dto.UpdatedAt),
		})
	}

	return result
}
