package api

import (
	"api_gateway_service/api/handler"
	"api_gateway_service/genproto"

	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
)

func NewRoutes(conn *grpc.ClientConn) *gin.Engine {
	mux := gin.Default()

	metro := genproto.NewMetroServiceClient(conn)
	user := genproto.NewUserServiceClient(conn)

	h := handler.NewHandler(metro, user)

	mux.POST("station/create", h.CreateStation)
	mux.GET("station/get:id", h.GetByIdStation)
	mux.GET("station/getall", h.GetAllStation)
	mux.PUT("station/update", h.UpdateStation)
	mux.DELETE("station/create", h.DeleteStation)

	mux.POST("card/create", h.CreateCard)
	mux.GET("card/get:id", h.GetByIdCard)
	mux.GET("card/getall", h.GetAll)
	mux.PUT("card/update", h.UpdateCard)
	mux.DELETE("card/create", h.DeleteCard)

	mux.POST("terminal/create", h.CreateTerminal)
	mux.GET("terminal/get:id", h.GetByIdTerminal)
	mux.GET("terminal/getall", h.GetAllTerminal)
	mux.PUT("terminal/update", h.UpdateTerminal)
	mux.DELETE("terminal/create", h.DeleteTerminal)

	mux.POST("transaction/create", h.CreateTransaction)
	mux.GET("transaction/get:id", h.GetByIdTransaction)
	mux.GET("transaction/getall", h.GetAllTransaction)
	mux.PUT("transaction/update", h.UpdateTransaction)
	mux.DELETE("transaction/create", h.DeleteTransction)

	return mux
}
