package usecase

import (
	"backend/apps/domain/auth/params"
	"backend/helper/response"
	"backend/pkg/encryption"
	"backend/pkg/token"
	"context"
	"database/sql"
)

// Login implements services.AuthSvc
func (a *authSvc) Login(ctx context.Context, req *params.UserLoginRequest) (*params.UserLoginResponse, *response.CustomError) {
	auth, err := a.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.NotFoundError()
		}

		return nil, response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	err = encryption.ValidatePassword(auth.Password, req.Password)
	if err != nil {
		return nil, response.NotFoundError()
	}

	token, err := token.GenerateToken(auth.Id)
	if err != nil {
		return nil, response.GeneralErrorWithAdditionalInfo(err.Error())
	}

	return &params.UserLoginResponse{
		Token: token,
	}, nil
}
