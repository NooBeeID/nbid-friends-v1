package params

import "backend/apps/domain/follow/models"

type GetMyFollowingResponse struct {
	Id     int    `json:"id"`
	Email  string `json:"email"`
	ImgUrl string `json:"img_url"`
}

func (g *GetMyFollowingResponse) AddModelToResponse(model *models.FollowWithAuth) {
	g.Email = model.Email
	g.Id = model.Id
	g.ImgUrl = model.ImgUrl
}
