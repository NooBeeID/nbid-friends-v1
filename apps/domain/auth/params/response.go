package params

import "backend/apps/domain/auth/models"

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserSearchResponse struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	ImgUrl   string `json:"img_url"`
	IsFollow bool   `json:"is_follow"`
}

func (u *UserSearchResponse) ParseModelToResponse(model *models.SearchAuth) {
	u.Email = model.Email
	u.Id = model.Id
	u.ImgUrl = model.ImgUrl
}

type UserProfileResponse struct {
	Id     int    `json:"id"`
	Email  string `json:"email"`
	ImgUrl string `json:"img_url"`
}

func (u *UserProfileResponse) ParseModelToResponse(model *models.Auth) {
	u.Email = model.Email
	u.Id = model.Id
	u.ImgUrl = model.ImgUrl
}
