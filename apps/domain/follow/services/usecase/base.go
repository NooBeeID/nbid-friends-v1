package usecase

import (
	"backend/apps/domain/follow/repositories"
	"backend/apps/domain/follow/services"
)

type followSvc struct {
	repo repositories.FollowRepo
}

func NewFollowSvc(repo repositories.FollowRepo) services.FollowSvc {
	return &followSvc{
		repo: repo,
	}
}
