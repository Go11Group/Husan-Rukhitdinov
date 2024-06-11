package model

type Users struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password_hash"`
	CreatedAt string `json:"created_at"`
}
