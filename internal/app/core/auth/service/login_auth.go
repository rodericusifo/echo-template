package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/rodericusifo/echo-template/internal/app/core/auth/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/app/core/auth/service/dto/output"
	"github.com/rodericusifo/echo-template/internal/pkg/types"

	mocks_pkg "github.com/rodericusifo/echo-template/mocks-pkg"
	pkg_types "github.com/rodericusifo/echo-template/pkg/types"
)

func (s *AuthService) LoginAuth(payload *input.LoginAuthDTO) (*output.LoginAuthDTO, error) {
	userModelRes, err := s.UserResource.GetUser(&pkg_types.Query{
		Searches: [][]pkg_types.SearchOperation{
			{
				{Field: "email", Operator: "=", Value: payload.Email},
			},
		},
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusNotFound, "user not found")
		}
		return nil, err
	}

	match := mocks_pkg.CheckHashPasswordUtil(payload.Password, userModelRes.Password)
	if !match {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "email and password not match")
	}

	claims := &types.JwtCustomClaims{
		XID:   userModelRes.XID,
		Name:  userModelRes.Name,
		Email: userModelRes.Email,
		Role:  userModelRes.Role,
	}

	token, err := mocks_pkg.GenerateJWTTokenFromClaimsUtil(claims)
	if err != nil {
		return nil, err
	}

	loginAuthDto := &output.LoginAuthDTO{
		Token: token,
	}

	return loginAuthDto, nil
}
