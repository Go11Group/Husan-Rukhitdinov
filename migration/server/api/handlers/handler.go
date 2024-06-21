package handlers

import "projects/storage/postgres"

type Handler struct {
	UserRepo *postgres.UserRepository
}

func NewHandler(us *postgres.UserRepository) *Handler {
	return &Handler{UserRepo: us}
}