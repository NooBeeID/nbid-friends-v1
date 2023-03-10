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
func (a *authSvc) Search(ctx context.Context, email string, id int) ([]*params.UserSearchResponse, *response.CustomError) {
	auth, err := a.repo.SearchAuthByEmail(ctx, email, id)
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
		tempAuthResp.IsFollow = false
		if search.AuthId != nil {
			if *search.AuthId == id {
				tempAuthResp.IsFollow = true
			}
		}
		authResp = append(authResp, tempAuthResp)
	}

	if authResp == nil {
		authResp = []*params.UserSearchResponse{}
	}

	return authResp, nil
}

// Profile implements services.AuthSvc
func (a *authSvc) Profile(ctx context.Context, authId int) (*params.UserProfileResponse, *response.CustomError) {
	auth, err := a.repo.FindById(ctx, authId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.NotFoundError()
		}

		return nil, response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	authResp := new(params.UserProfileResponse)

	authResp.ParseModelToResponse(auth)

	return authResp, nil
}
