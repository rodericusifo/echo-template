package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/rodericusifo/echo-template/internal/app/core/auth/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/app/model/database/sql"
	"github.com/rodericusifo/echo-template/pkg/types"

	mocks_pkg "github.com/rodericusifo/echo-template/mocks-pkg"
)

func (s *AuthService) RegisterAuth(payload *input.RegisterAuthDTO) error {
	userModelRes, err := s.UserResource.GetUser(&types.Query{
		Selects: []types.SelectOperation{
			{Field: "id"},
		},
		Searches: [][]types.SearchOperation{
			{
				{Field: "email", Operator: "=", Value: payload.Email},
			},
		},
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if userModelRes != nil {
		return echo.NewHTTPError(http.StatusConflict, "user already registered")
	}

	hashedPassword, err := mocks_pkg.GenerateHashPasswordUtil(payload.Password)
	if err != nil {
		return err
	}

	userModel := &sql.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashedPassword,
		Role:     payload.Role,
	}
	err = s.UserResource.CreateUser(userModel)
	if err != nil {
		return err
	}

	return nil
}
