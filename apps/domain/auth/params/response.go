package params

import "backend/apps/domain/auth/models"

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserSearchResponse struct {
	Id     int    `json:"id"`
	Email  string `json:"email"`
	ImgUrl string `json:"img_url"`
}

func (u *UserSearchResponse) ParseModelToResponse(model *models.Auth) {
	u.Email = model.Email
	u.Id = model.Id
	u.ImgUrl = model.ImgUrl
}
