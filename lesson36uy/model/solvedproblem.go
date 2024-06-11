package model

type SolvedProblem struct {
	ID        string
	UserID    string `json:"user_id"`
	ProblemID string `json:"problem_id"`
	SolvedAt  string `json:"solved_at"`
}
