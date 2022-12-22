package usecase

import (
	"backend/apps/commons/response"
	"backend/apps/domain/auth/params"
	"backend/pkg/encryption"
	"context"
)

// Register implements services.AuthSvc
func (a *authSvc) Register(ctx context.Context, req *params.UserRegisterRequest) *response.CustomError {
	reqModel := req.ParseToModel()

	hashPass, err := encryption.HashPassword(reqModel.Password)
	if err != nil {
		return response.GeneralError()
	}

	reqModel.Password = hashPass

	err = a.repo.Create(ctx, reqModel)
	if err != nil {
		return response.RepositoryErrorWithAdditionalInfo(err.Error())
	}
	return nil
}
