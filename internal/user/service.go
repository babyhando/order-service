package user

import (
	"context"
	"errors"
	"log"
	"order-service/internal/user/domain"
	"order-service/internal/user/port"
)

var (
	ErrUserOnCreate = errors.New("error on creating new user")
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(ctx context.Context, user domain.User) (domain.UserID, error) {
	userID, err := s.repo.Create(ctx, user)
	if err != nil {
		log.Println("error on creating new user : ", err.Error())
		return 0, ErrUserOnCreate
	}

	return userID, nil
}

func (s *service) GetUserByID(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	panic("not implemented")
}