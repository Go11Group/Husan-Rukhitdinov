package service

import (

	pb "google.golang.org/grpc/metro_service/genproto/metro_service"
	"google.golang.org/grpc/metro_service/storage/postgres"
)

type cardService struct {
	pb.UnimplementedMetroServiceServer
	C *postgres.CardRepo
}

func NewCardService(C *postgres.CardRepo) *cardService {
	return &cardService{C: C}
}

