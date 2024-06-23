package handler

import (
	"database/sql"
	"user_pro/storage/postgres"
)

type Handler struct {
	db   *sql.DB
	User *postgres.UserRepo
}

func NewHandler(db *sql.DB, user *postgres.UserRepo) *Handler {
	return &Handler{
		db:   db,
		User: user,
	}
}
