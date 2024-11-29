package storage

import (
	"context"
	"errors"
	"order-service/internal/user/domain"
	"order-service/internal/user/port"
	"order-service/pkg/adapters/storage/mapper"
	"order-service/pkg/adapters/storage/types"
	"order-service/pkg/cache"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB, cached bool, provider cache.Provider) port.Repo {
	repo := &userRepo{db}
	if !cached {
		return repo
	}

	return &userCachedRepo{
		repo:     repo,
		provider: provider,
	}
}

func (r *userRepo) Create(ctx context.Context, userDomain domain.User) (domain.UserID, error) {
	user := mapper.UserDomain2Storage(userDomain)
	return domain.UserID(user.ID), r.db.Table("users").WithContext(ctx).Create(user).Error
}

func (r *userRepo) GetByID(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	var user types.User
	err := r.db.Table("users").
		Where("id = ?", userID).
		First(&user).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if user.ID == 0 {
		return nil, nil
	}

	return mapper.UserStorage2Domain(user), nil
}
