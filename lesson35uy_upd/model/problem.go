package model

import "github.com/google/uuid"

type Problem struct {
	ProblemID   uuid.UUID `json:"problem_id"`
	Title       string    `json:"title"`
	Difficulty  string    `json:"difficulty"`
	Description string    `json:"description"`
}
