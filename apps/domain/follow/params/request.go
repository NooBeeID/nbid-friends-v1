package params

import "backend/apps/domain/follow/models"

type FollowingRequest struct {
	FollowingId int `json:"following_id"`
	AuthId      int `json:"auth_id"`
}

func (f *FollowingRequest) ParseToModel() *models.Follow {
	return &models.Follow{
		AuthId:      f.AuthId,
		FollowingId: f.FollowingId,
	}
}
