package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/rodericusifo/echo-template/internal/pkg/constant"
	"github.com/rodericusifo/echo-template/pkg/util"

	internal_pkg_util "github.com/rodericusifo/echo-template/internal/pkg/util"
)

func UserRolesPermission(roles ...constant.UserRole) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqUser := internal_pkg_util.GetRequestUser(c)

			if reqUser.Role == constant.UserRole("") {
				return echo.NewHTTPError(echo.ErrUnprocessableEntity.Code, "user role not exist")
			}

			if len(roles) > 0 {
				allowed := util.CheckSliceContain(roles, reqUser.Role)
				if !allowed {
					return echo.NewHTTPError(echo.ErrUnauthorized.Code, fmt.Sprint("user role not allowed. allowed roles:", roles))
				}
			}

			return next(c)
		}
	}
}
