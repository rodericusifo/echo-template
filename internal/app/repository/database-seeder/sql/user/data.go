package user

import (
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
)

var UserSeedData = []*UserSeedPayload{
	// ADMIN
	{
		XID:      "8ea778bc-3958-4e9f-8fa2-a8a9ad8f2ab1",
		Name:     "admin",
		Email:    "admin@gmail.com",
		Password: "p4ssw0rd",
		Role:     constant.ADMIN,
	},
	{
		XID:      "c14c4865-428a-410d-a20d-6c90feb19d77",
		Name:     "admin.one",
		Email:    "admin.one@gmail.com",
		Password: "p4ssw0rd",
		Role:     constant.ADMIN,
	},
}
