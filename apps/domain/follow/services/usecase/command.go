package usecase

import (
	"backend/apps/commons/response"
	"backend/apps/domain/follow/params"
	"context"
	"time"
)

// FollowFriend implements services.FollowSvc
func (f *followSvc) FollowFriend(ctx context.Context, req *params.FollowingRequest) *response.CustomError {
	follow := req.ParseToModel()
	follow.CreatedAt = time.Now()

	err := f.repo.Create(ctx, follow)
	if err != nil {
		return response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	return nil
}

// UnfollowFriend implements services.FollowSvc
func (f *followSvc) UnfollowFriend(ctx context.Context, req *params.UnfollRequest) *response.CustomError {
	unfoll := req.ParseToModel()

	err := f.repo.Delete(ctx, unfoll.AuthId, unfoll.FollowingId)
	if err != nil {
		return response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	return nil

}
