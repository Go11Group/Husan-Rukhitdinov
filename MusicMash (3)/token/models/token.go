package models

type TokenReq struct {
	UserId   string
	UserName string
	Email    string
}

type TokenRes struct {
	Token string `json:"token"`
}
