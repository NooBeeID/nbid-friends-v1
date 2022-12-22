package usecase

import (
	"backend/apps/commons/response"
	"backend/apps/domain/follow/params"
	"context"
	"database/sql"
)

// GetMyFollowing implements services.FollowSvc
func (f *followSvc) GetMyFollowing(ctx context.Context, authId int) ([]*params.GetMyFollowingResponse, *response.CustomError) {
	following, err := f.repo.GetAll(ctx, authId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.NotFoundError()
		}
		return nil, response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	var followingResponse []*params.GetMyFollowingResponse

	for _, follow := range following {
		followingResp := params.GetMyFollowingResponse{}
		followingResp.AddModelToResponse(follow)

		followingResponse = append(followingResponse, &followingResp)
	}

	if followingResponse == nil {
		followingResponse = []*params.GetMyFollowingResponse{}
	}

	return followingResponse, nil

}
