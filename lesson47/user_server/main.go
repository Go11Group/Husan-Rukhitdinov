package main

import (
	"net"

	pb "github.com/Go11Group/Husan-Rukhitdinov/lesson44/user_server/genproto/user_service"
	"github.com/Go11Group/Husan-Rukhitdinov/lesson44/user_server/storage/postgres"
	"github.com/Go11Group/Husan-Rukhitdinov/lesson44/user_server/service"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	listener, err := net.Listen("tcp", ":8088")
	if err != nil {
		panic(err)
	}

	user := service.NewUserService(postgres.NewUserRepo(db))

	u := grpc.NewServer()
	pb.RegisterUserServiceServer(u, user)

	err = u.Serve(listener)
	if err != nil {
		panic(err)
	}
}
