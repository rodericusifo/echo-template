package service

import (
	"github.com/rodericusifo/echo-template/internal/app/core/auth/service/dto/input"
	"github.com/rodericusifo/echo-template/internal/app/core/auth/service/dto/output"
	"github.com/rodericusifo/echo-template/internal/app/core/user/resource"
)

type IAuthService interface {
	LoginAuth(payload *input.LoginAuthDTO) (*output.LoginAuthDTO, error)
}

type AuthService struct {
	UserResource resource.IUserResource
}

func InitAuthService(userResource resource.IUserResource) IAuthService {
	return &AuthService{
		UserResource: userResource,
	}
}
