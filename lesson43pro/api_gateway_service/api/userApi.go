package api

import (
	"net/http"

	"github.com/Go11Group/at_lesson/lesson43/api_gateway_service/api/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes() *http.Server {
	uurl := gin.Default()

	l := handler.NewHandler()

	uurl.POST("station/create", l.Client)
	uurl.GET("station/get:id", l.Client)
	uurl.GET("station/getall", l.Client)
	uurl.PUT("station/update", l.Client)
	uurl.DELETE("station/create", l.Client)

	return &http.Server{Handler: uurl, Addr: ":8090"}

}
