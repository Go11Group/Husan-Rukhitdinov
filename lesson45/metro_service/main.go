package main

import (
	"net"

	pbCard "google.golang.org/grpc/metro_service/genproto/service"
	pbStation "google.golang.org/grpc/metro_service/genproto/service"
	pbTerminal "google.golang.org/grpc/metro_service/genproto/service"
	pbTransaction "google.golang.org/grpc/metro_service/genproto/service"
	"google.golang.org/grpc/metro_service/service"
	"google.golang.org/grpc/metro_service/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	cardService := service.NewCardService(postgres.NewCardRepo(db))
	stationService := service.NewStationService(postgres.NewStationRepo(db))
	terminalService := service.NewTerminalService(postgres.NewTerminalRepo(db))
	transactionService := service.NewTransactionService(postgres.NewTransactionRepo(db))

	grpcServer := grpc.NewServer()
	pbCard.RegisterCardServiceServer(grpcServer, cardService)
	pbStation.RegisterStationServiceServer(grpcServer, stationService)
	pbTerminal.RegisterTerminalServiceServer(grpcServer, terminalService)
	pbTransaction.RegisterTransactionServiceServer(grpcServer, transactionService)

	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
