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

type Problem struct {
	ProblemID   uuid.UUID    `json:"problem_id"`
	Title       string `json:"title"`
	Difficulty  string `json:"difficulty"`
	Description string `json:"description"`
}

type SolvedProblem struct {
	ID        uuid.UUID
	UserID    uuid.NullUUID       `json:"user_id"`
	ProblemID uuid.NullUUID       `json:"problem_id"`
	SolvedAt  time.Time `json:"solved_at"`
}
