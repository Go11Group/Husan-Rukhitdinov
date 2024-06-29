package handler

import "api_gateway_service/genproto"

type Handler struct {
	Metro *genproto.NewMetroServiceClient
	User  *genproto.NewUserServiceClient
}

func NewHandler(Metro genproto.NewMetroServiceClient, User genproto.NewUserServiceClient) *Handler {
	return &Handler{Metro: Metro, User: User}
}
