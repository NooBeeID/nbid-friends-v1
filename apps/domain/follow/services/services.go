package services

import (
	"backend/apps/commons/response"
	"backend/apps/domain/follow/params"
	"context"
)

type FollowSvc interface {
	FollowFriend(ctx context.Context, req *params.FollowingRequest) *response.CustomError
	GetMyFollowing(ctx context.Context, authId int) ([]*params.GetMyFollowingResponse, *response.CustomError)
	UnfollowFriend(ctx context.Context, req *params.UnfollRequest) *response.CustomError
}
