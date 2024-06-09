package model

import "time"

type Users struct {
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password_hash"`
	CreatedAt time.Time `json:"created_at"`
}

type Problem struct {
	ProblemID   int    `json:"problem_id"`
	Title       string `json:"title"`
	Difficulty  string `json:"difficulty"`
	Description string `json:"description"`
}

type SolvedProblem struct {
	ID        int
	UserID    int       `json:"user_id"`
	ProblemID int       `json:"problem_id"`
	SolvedAt  time.Time `json:"solved_at"`
}
