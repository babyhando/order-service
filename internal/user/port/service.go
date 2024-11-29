package port

import (
	"context"
	"order-service/internal/user/domain"
)

type Service interface {
	CreateUser(ctx context.Context, user domain.User) (domain.UserID, error)
	GetUserByID(ctx context.Context, userID domain.UserID) (*domain.User, error)
}
