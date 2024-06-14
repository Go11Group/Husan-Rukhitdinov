package models

import "time"

type Courses struct {
	ID          string     `json:"course_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdateAt    time.Time  `json:"update_at"`
	DeleteAt    *time.Time `json:"deleted_at"`
}
type CourseF struct {
	ID          string
	Title       string
	Description string
}