package usecase

import (
	"backend/apps/commons/response"
	"backend/apps/domain/auth/params"
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

// Search implements services.AuthSvc
func (a *authSvc) Search(ctx context.Context, email string) ([]*params.UserSearchResponse, *response.CustomError) {
	auth, err := a.repo.SearchAuthByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.NotFoundError()
		}

		return nil, response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	var authResp []*params.UserSearchResponse
	for _, search := range auth {
		tempAuthResp := new(params.UserSearchResponse)
		tempAuthResp.ParseModelToResponse(search)

		authResp = append(authResp, tempAuthResp)
	}

	return authResp, nil
}
