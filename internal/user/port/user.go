package port

import (
	"context"
	"order-service/internal/user/domain"
)

type Repo interface {
	Create(ctx context.Context, user domain.User) (domain.UserID, error)
	GetByID(ctx context.Context, userID domain.UserID) (*domain.User, error)
}
