package models

import (
	"time"
)

type Users struct {
	ID        string    `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Birthday  string    `json:"birthday"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	DeleteAt  int       `json:"deleted_at"`
}

type UserF struct {
	ID       string
	Name     string
	Email    string
	Birthday string
	Password string
}
