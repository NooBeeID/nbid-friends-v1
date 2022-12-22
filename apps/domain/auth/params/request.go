package params

import "backend/apps/domain/auth/models"

type UserRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	ImgUrl   string `json:"img_url"`
}

func (u *UserRegisterRequest) ParseToModel() *models.Auth {
	return &models.Auth{
		Email:    u.Email,
		Password: u.Password,
		ImgUrl:   u.ImgUrl,
	}
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSearchRequest struct {
	Email string `json:"email"`
}
