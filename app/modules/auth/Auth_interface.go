package auth

import (
	authDtos "fotongo/app/modules/auth/dtos"
	"fotongo/app/utils/dtos"
)

type AuthInterface interface {
	Login(request authDtos.LoginRequest) (dtos.JSONResponses, error)
	Register(request authDtos.RegisterRequest) (dtos.JSONResponses, error)
}
