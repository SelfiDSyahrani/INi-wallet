package dto

import "dev_selfi/model"

type UserRequestBody struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRequestParams struct {
	UserWallet_ID string `uri:"id" binding:"required"`
}

type UserRequestQuery struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

type UserResponseBody struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FormatUser(user *model.User) UserResponseBody {
	formattedUser := UserResponseBody{}
	formattedUser.ID = user.UserWallet_ID
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
