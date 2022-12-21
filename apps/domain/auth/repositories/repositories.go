package repositories

import (
	"backend/apps/domain/auth/models"
	"context"
)

type AuthRepo interface {
	Create(ctx context.Context, auth *models.Auth) error
	FindByEmail(ctx context.Context, email string) (*models.Auth, error)
}
