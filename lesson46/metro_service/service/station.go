package service

import (
	pb "google.golang.org/grpc/metro_service/genproto/metro_service"
	"google.golang.org/grpc/metro_service/storage/postgres"
)

type stationService struct {
	pb.UnimplementedMetroServiceServer
	S *postgres.StationRepo
}

func NewStationService(S *postgres.StationRepo) *stationService {
	return &stationService{S: S}
}