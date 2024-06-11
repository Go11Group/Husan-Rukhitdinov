package model

type Problem struct {
	ProblemID   string `json:"problem_id"`
	Title       string `json:"title"`
	Difficulty  string `json:"difficulty"`
	Description string `json:"description"`
}
