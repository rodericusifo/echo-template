package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/internal/pkg/request"
	"github.com/rodericusifo/echo-template/internal/pkg/types"
	"github.com/rodericusifo/echo-template/internal/pkg/validator"

	mocks_pkg "github.com/rodericusifo/echo-template/mocks-pkg"
	pkg_types "github.com/rodericusifo/echo-template/pkg/types"
)

func UserRequest() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := c.Get(constant.C_KEY_USER).(*jwt.Token).Claims
			user, ok := claims.(*types.JwtCustomClaims)
			if !ok {
				return echo.NewHTTPError(echo.ErrUnprocessableEntity.Code, fmt.Sprintf("invalid claims type. correct type: %T", claims))
			}

			userModelRes, err := mocks_pkg.UserResource().GetUser(&pkg_types.Query{
				Selects: []pkg_types.SelectOperation{
					{Field: "id"},
				},
				Searches: [][]pkg_types.SearchOperation{
					{
						{Field: "xid", Operator: "=", Value: user.XID},
					},
				},
			})
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return echo.NewHTTPError(http.StatusNotFound, "user not found")
				}
				return err
			}

			reqUser := new(request.RequestUser)
			reqUser = &request.RequestUser{
				ID:    userModelRes.ID,
				XID:   user.XID,
				Name:  user.Name,
				Email: user.Email,
				Role:  user.Role,
			}
			if err := validator.ValidateRequestUser(reqUser); err != nil {
				return err
			}

			c.Set(constant.C_KEY_REQUEST_USER, reqUser)
			return next(c)
		}
	}
}
