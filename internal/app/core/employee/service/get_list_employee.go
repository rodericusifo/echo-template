package service

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/app/core/employee/service/dto/output"
	"github.com/rodericusifo/echo-template/internal/pkg/util/mapper"
	"github.com/rodericusifo/echo-template/pkg/response/structs"
	"github.com/rodericusifo/echo-template/pkg/types"
	"github.com/rodericusifo/echo-template/pkg/util"
)

func (s *EmployeeService) GetListEmployee(payload *input.GetListEmployeeDTO) ([]*output.EmployeeDTO, *structs.Meta, error) {
	page, limit := util.DefinePageLimitPagination(payload.Page, payload.Limit)

	employeeListModelRes, countEmployeeListModelRes, err := s.EmployeeResource.GetListEmployeeAndCount(&types.Query{
		Offset: util.CountOffsetPagination(page, limit),
		Limit:  limit,
		Searches: [][]types.SearchOperation{
			{
				{Field: "user_id", Operator: "=", Value: payload.UserID},
			},
		},
	})
	if err != nil {
		return nil, nil, err
	}

	if len(employeeListModelRes) < 1 {
		return nil, nil, echo.NewHTTPError(http.StatusNotFound, "list employee not found")
	}

	countEmployeeAllModelRes, err := s.EmployeeResource.CountAllEmployee(&types.Query{
		Searches: [][]types.SearchOperation{
			{
				{Field: "user_id", Operator: "=", Value: payload.UserID},
			},
		},
	})
	if err != nil {
		return nil, nil, err
	}

	employeeListDto := mapper.MapEmployeesToEmployeeDTOs(employeeListModelRes)

	meta := &structs.Meta{
		CurrentPage:      page,
		CountDataPerPage: countEmployeeListModelRes,
		TotalData:        countEmployeeAllModelRes,
	}

	meta.TotalPage = util.CountTotalPagePagination(meta.CountDataPerPage, meta.TotalData)

	return employeeListDto, meta, nil
}
