package storage

import (
	"context"
	"errors"

	"github.com/babyhando/order-service/internal/user/domain"
	"github.com/babyhando/order-service/internal/user/port"
	"github.com/babyhando/order-service/pkg/adapters/storage/mapper"
	"github.com/babyhando/order-service/pkg/adapters/storage/types"
	"github.com/babyhando/order-service/pkg/cache"

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

func (r *userRepo) GetByFilter(ctx context.Context, filter *domain.UserFilter) (*domain.User, error) {
	var user types.User

	q := r.db.Table("users").Debug().WithContext(ctx)

	if filter.ID > 0 {
		q = q.Where("id = ?", filter.ID)
	}

	if len(filter.Phone) > 0 {
		q = q.Where("phone = ?", filter.Phone)
	}

	err := q.First(&user).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if user.ID == 0 {
		return nil, nil
	}

	return mapper.UserStorage2Domain(user), nil
}
