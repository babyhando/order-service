package storage

import (
	"context"
	"errors"

	"github.com/babyhando/order-service/internal/order/domain"
	orderPort "github.com/babyhando/order-service/internal/order/port"
	"github.com/babyhando/order-service/pkg/adapters/storage/mapper"
	"github.com/babyhando/order-service/pkg/adapters/storage/types"

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
	o := mapper.OrderDomain2Storage(order)
	return domain.OrderID(o.ID), r.db.Table("orders").WithContext(ctx).Create(o).Error
}

func (r *orderRepo) GetByID(ctx context.Context, orderID domain.OrderID) (*domain.Order, error) {
	var order types.Order

	err := r.db.Table("orders").WithContext(ctx).Where("id = ?", orderID).First(&order).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if order.ID == 0 {
		return nil, nil
	}

	return mapper.OrderStorage2Domain(order)
}

func (r *orderRepo) Get(ctx context.Context, filter domain.OrderListFilters) ([]domain.Order, error) {
	panic("not implemented")
}
