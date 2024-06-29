package service

import (
	pb "google.golang.org/grpc/metro_service/genproto/metro_service"
	"google.golang.org/grpc/metro_service/storage/postgres"
)

type transactionService struct {
	pb.UnimplementedMetroServiceServer
	Tr *postgres.TransactionRepo
}

func NewTransactionService(Tr *postgres.TransactionRepo) *transactionService {
	return &transactionService{Tr: Tr}
}
