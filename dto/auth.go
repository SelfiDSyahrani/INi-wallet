package dto

import "dev_selfi/model"

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseBody struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatLogin(user *model.User, token string) LoginResponseBody {
	return LoginResponseBody{
		ID:    user.UserWallet_ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
}
