package storage

import (
	"context"
	"order-service/internal/order/domain"
	orderPort "order-service/internal/order/port"

	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) orderPort.Repo {
	return &orderRepo{
		db: db,
	}
}

func (r *orderRepo) Create(ctx context.Context, order domain.Order) (domain.OrderID, error) {
	panic("not implemented")
}

func (r *orderRepo) GetByID(ctx context.Context, orderID domain.OrderID) (*domain.Order, error) {
	panic("not implemented")
}

func (r *orderRepo) Get(ctx context.Context, filter domain.OrderListFilters) ([]domain.Order, error) {
	panic("not implemented")
}
