package models

type RegisterReq struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}
