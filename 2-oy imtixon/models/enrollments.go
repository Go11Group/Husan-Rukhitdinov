package models

import "time"

type Enrollments struct {
	ID             string    `json:"enrollment_id"`
	UserID         string    `json:"user_id"`
	CourseID       string    `json:"course_id"`
	EnrollmentDate string    `json:"enrollment_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdateAt       time.Time `json:"update_at"`
	DeleteAt       int       `json:"deleted_at"`
}
type EnrollmentF struct {
	ID             string
	UserID         string
	CourseID       string
	EnrollmentDate string
}
