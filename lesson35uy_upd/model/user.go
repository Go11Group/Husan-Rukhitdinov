package model

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	UserID    uuid.UUID `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password_hash"`
	CreatedAt time.Time `json:"created_at"`
}
