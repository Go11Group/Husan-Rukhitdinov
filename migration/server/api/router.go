package api

import (
	"net/http"
	"projects/api/handlers"
	"projects/storage/postgres"
)

func RooterApi(user *postgres.UserRepository) *http.ServeMux {
	handler := handlers.NewHandler(user)
	router := http.NewServeMux()
	router.HandleFunc("POST /user/create", handler.CreateUser)
	router.HandleFunc("DELETE /user/delete", handler.DeleteUser)
	router.HandleFunc("UPDATE /user/update/", handler.UpdateUser)
	router.HandleFunc("GET  /user/get", handler.GetUser)
	return router

}
