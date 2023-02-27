package dto

import (
	"INi-Wallet/model"
	"time"
)

type UserRequestBody struct {
	Name     string  `json:"name" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Phone    string  `json:"phone" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Balance  float64 `json:"balance" binding:"required"`
}

type UserRequestParams struct {
	UserWallet_ID string `uri:"id" binding:"required"`
}

type UserRequestQuery struct {
	Name string `form:"name"`
}

type UserResponseBody struct {
	UserWallet_ID string    `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseBody struct {
	UserWallet_ID string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Token         string `json:"token"`
}

func FormatUser(user *model.User) UserResponseBody {
	formattedUser := UserResponseBody{}
	formattedUser.UserWallet_ID = user.UserWallet_ID
	formattedUser.Name = user.Name
	formattedUser.Email = user.Email
	return formattedUser
}

func FormatUsers(authors []*model.User) []UserResponseBody {
	formattedUsers := []UserResponseBody{}
	for _, user := range authors {
		formattedUser := FormatUser(user)
		formattedUsers = append(formattedUsers, formattedUser)
	}
	return formattedUsers
}

func FormatLogin(user *model.User, token string) LoginResponseBody {
	return LoginResponseBody{
		UserWallet_ID:    user.UserWallet_ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
}
