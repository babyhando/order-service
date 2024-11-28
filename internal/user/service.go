package user

import (
	"context"
	"order-service/internal/user/domain"
	"order-service/internal/user/port"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetUserByID(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	panic("not implemented")
}
