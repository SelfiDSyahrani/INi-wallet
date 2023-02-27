package model

type User struct {
	UserWallet_ID string `json:"userWallet_ID"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	Balance float64 `json:"balance"`
}