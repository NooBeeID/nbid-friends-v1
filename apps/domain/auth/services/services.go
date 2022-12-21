package services

import (
	"backend/apps/domain/auth/params"
	"backend/helper/response"
	"context"
)

type AuthSvc interface {
	Register(ctx context.Context, req *params.UserRegisterRequest) *response.CustomError
	Login(ctx context.Context, req *params.UserLoginRequest) (*params.UserLoginResponse, *response.CustomError)
}
