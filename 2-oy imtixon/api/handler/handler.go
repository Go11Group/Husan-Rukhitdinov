package handler

import (
	"database/sql"
	"project/storage/postgres"
)

type Handler struct {
	db         *sql.DB
	User       *postgres.User
	Course     *postgres.Courses
	Lesson     *postgres.Lessons
	Enrollment *postgres.Enrollments
}

func NewHandler(db *sql.DB, user *postgres.User, course *postgres.Courses, lesson *postgres.Lessons, enrollment *postgres.Enrollments) *Handler {
	return &Handler{
		db:     db,
		User:   user,
		Course: course,
		Lesson: lesson,
		Enrollment: enrollment,
	}
}
