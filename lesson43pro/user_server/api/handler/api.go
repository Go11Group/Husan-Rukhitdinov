package handler

import (
	"database/sql"
	"net/http"
	// "user_pro/api/handler"
	"user_pro/storage/postgres"

	"github.com/gin-gonic/gin"
)

func UserRoutes(db *sql.DB, user *postgres.UserRepo) *http.Server {
	url := gin.Default()

	l := NewHandler(db, user)

	url.POST("station/create", l.CreateUser)
	url.GET("station/get:id", l.GetByIdUser)
	url.GET("station/getall", l.GetAllUser)
	url.PUT("station/update", l.UpdateUser)
	url.DELETE("station/create", l.DeleteUser)

	return &http.Server{Handler: url, Addr: "8090"}
}
