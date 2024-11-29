package storage

import (
	"context"
	"order-service/internal/user/domain"
	"order-service/internal/user/port"
	"order-service/pkg/adapters/storage/mapper"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) port.Repo {
	return &userRepo{db}
}

func (r *userRepo) Create(ctx context.Context, userDomain domain.User) (domain.UserID, error) {
	user := mapper.UserDomain2Storage(userDomain)
	return domain.UserID(user.ID), r.db.Table("users").WithContext(ctx).Create(user).Error
}
