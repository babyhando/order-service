package port

import (
	"context"
	"order-service/internal/order/domain"
	"time"
)

type Service interface {
	CreateOrder(ctx context.Context, order domain.Order) (domain.OrderID, error)
	SubmitOrder(ctx context.Context, orderID domain.OrderID, submittedAt time.Time) error
}
