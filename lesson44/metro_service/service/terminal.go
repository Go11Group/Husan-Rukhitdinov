package service

import (
	pb "google.golang.org/grpc/metro_service/genproto/metro_service"
	"google.golang.org/grpc/metro_service/storage/postgres"
)

type terminalService struct {
	pb.UnimplementedMetroServiceServer
	T *postgres.TerminalRepo
}

func NewTerminalService(T *postgres.TerminalRepo) *terminalService {
	return &terminalService{T: T}
}