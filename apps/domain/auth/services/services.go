package services

import (
	"backend/apps/commons/response"
	"backend/apps/domain/auth/params"
	"context"
)

type AuthSvc interface {
	Register(ctx context.Context, req *params.UserRegisterRequest) *response.CustomError
	Login(ctx context.Context, req *params.UserLoginRequest) (*params.UserLoginResponse, *response.CustomError)
	Search(ctx context.Context, email string) ([]*params.UserSearchResponse, *response.CustomError)
}
