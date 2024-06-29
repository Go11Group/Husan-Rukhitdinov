package main

import (
	"log"

	"google.golang.org/grpc"
	"api_gateway_service/api"
	"api_gateway_service/genproto"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	metro := genproto.NewMetroServiceClient(conn)
	user := genproto.NewUserServiceClient(conn)

	r := api.NewRoutes(conn, metro, user)

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
