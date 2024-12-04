package port

import (
	"context"

	"github.com/babyhando/order-service/internal/order/domain"
)

type Repo interface {
	Create(ctx context.Context, order domain.Order) (domain.OrderID, error)
	GetByID(ctx context.Context, orderID domain.OrderID) (*domain.Order, error)
	Get(ctx context.Context, filter domain.OrderListFilters) ([]domain.Order, error)
}
