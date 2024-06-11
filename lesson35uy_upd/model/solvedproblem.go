package model

import (
	"time"

	"github.com/google/uuid"
)


type SolvedProblem struct {
	ID        uuid.UUID
	UserID    uuid.NullUUID       `json:"user_id"`
	ProblemID uuid.NullUUID       `json:"problem_id"`
	SolvedAt  time.Time `json:"solved_at"`
}
