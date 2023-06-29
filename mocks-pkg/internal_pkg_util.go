package mocks_pkg

import (
	"github.com/rodericusifo/echo-template/internal/pkg/util"
)

var (
	GenerateHashPasswordUtil       = util.GenerateHashPassword
	GenerateJWTTokenFromClaimsUtil = util.GenerateJWTTokenFromClaims

	CheckHashPasswordUtil = util.CheckHashPassword
)
