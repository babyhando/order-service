package port

import (
	"context"
	"order-service/internal/user/domain"
)

type Service interface {
	GetUserByID(ctx context.Context, userID domain.UserID) (*domain.User, error)
}
