package usecase

import (
	"backend/apps/domain/auth/repositories"
	"backend/apps/domain/auth/services"
)

type authSvc struct {
	repo repositories.AuthRepo
}

func NewAuthSvc(repo repositories.AuthRepo) services.AuthSvc {
	return &authSvc{
		repo: repo,
	}
}
