package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"order-service/internal/user/domain"
	"order-service/internal/user/port"
)

var (
	ErrUserOnCreate           = errors.New("error on creating new user")
	ErrUserCreationValidation = errors.New("validation failed")
	ErrUserNotFound           = errors.New("user not found")
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
	if err := user.Validate(); err != nil {
		return 0, fmt.Errorf("%w %w", ErrUserCreationValidation, err)
	}

	userID, err := s.repo.Create(ctx, user)
	if err != nil {
		log.Println("error on creating new user : ", err.Error())
		return 0, ErrUserOnCreate
	}

	return userID, nil
}

func (s *service) GetUserByFilter(ctx context.Context, filter *domain.UserFilter) (*domain.User, error) {
	user, err := s.repo.GetByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}
