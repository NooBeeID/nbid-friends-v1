package repositories

import (
	"backend/apps/domain/follow/models"
	"context"
)

type FollowRepo interface {
	Create(ctx context.Context, follow *models.Follow) error
	GetAll(ctx context.Context, authId int) ([]*models.FollowWithAuth, error)
	Delete(ctx context.Context, authId, followingId int) error
}
