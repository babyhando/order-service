package port

import (
	"context"
	"time"

	"github.com/babyhando/order-service/internal/order/domain"
)

type Service interface {
	CreateOrder(ctx context.Context, order domain.Order) (domain.OrderID, error)
	SubmitOrder(ctx context.Context, orderID domain.OrderID, submittedAt time.Time) error
}
