package service

import (
	"context"
	"order-service/api/pb"
	userPort "order-service/internal/user/port"
)

type UserService struct {
	svc userPort.Service
}

func (s *UserService) SignUp(ctx context.Context, req *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {
	panic("not implemented")
}
