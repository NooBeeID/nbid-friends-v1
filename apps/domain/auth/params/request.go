package params

import "backend/apps/domain/auth/models"

type UserRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserRegisterRequest) ParseToModel() *models.Auth {
	return &models.Auth{
		Email:    u.Email,
		Password: u.Password,
	}
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
