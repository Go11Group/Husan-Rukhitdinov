package api

import (
	"github.com/Go11Group/at_lesson/lesson43/api_gateway_service/api/handler"

	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes() *http.Server {
	mux := gin.Default()

	h := handler.NewHandler()

	// mux.POST("/station/create", h.Client)

	mux.POST("station/create", h.Client)
	mux.GET("station/get:id", h.Client)
	mux.GET("station/getall", h.Client)
	mux.PUT("station/update", h.Client)
	mux.DELETE("station/create", h.Client)

	mux.POST("card/create", h.Client)
	mux.GET("card/get:id", h.Client)
	mux.GET("card/getall", h.Client)
	mux.PUT("card/update", h.Client)
	mux.DELETE("card/create", h.Client)

	mux.POST("terminal/create", h.Client)
	mux.GET("terminal/get:id", h.Client)
	mux.GET("terminal/getall", h.Client)
	mux.PUT("terminal/update", h.Client)
	mux.DELETE("terminal/create", h.Client)

	mux.POST("transaction/create", h.Client)
	mux.GET("transaction/get:id", h.Client)
	mux.GET("transaction/getall", h.Client)
	mux.PUT("transaction/update", h.Client)
	mux.DELETE("transaction/create", h.Client)

	return &http.Server{Handler: mux, Addr: ":8081"}
}
