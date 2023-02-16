package types

import (
	"github.com/golang-jwt/jwt/v5"

	"github.com/rodericusifo/echo-template/internal/pkg/constant"
)

type JwtCustomClaims struct {
	XID   string            `json:"xid"`
	Name  string            `json:"name"`
	Email string            `json:"email"`
	Role  constant.UserRole `json:"role"`
	jwt.RegisteredClaims
}
