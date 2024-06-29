package service

import (
	pb "github.com/Go11Group/Husan-Rukhitdinov/lesson44/user_server/genproto/user_service"
	"github.com/Go11Group/Husan-Rukhitdinov/lesson44/user_server/storage/postgres"
)

type userService struct {
	pb.UnimplementedUserServiceServer
	C *postgres.UserRepo
}

func NewUserService(C *postgres.UserRepo) *userService {
	return &userService{C: C}
}
