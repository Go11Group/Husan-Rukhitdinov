package models

import "time"

type Lessons struct {
	ID        string    `json:"lesson_id"`
	CourseID  string    `json:"course_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	DeleteAt  string    `json:"deleted_at"`
}
type LessonF struct {
	ID       string
	CourseID string
	Title    string
	Content  string
}
