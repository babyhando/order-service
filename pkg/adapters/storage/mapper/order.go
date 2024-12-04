package mapper

import (
	"github.com/babyhando/order-service/internal/order/domain"
	userDomain "github.com/babyhando/order-service/internal/user/domain"
	"github.com/babyhando/order-service/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

func OrderDomain2Storage(orderDomain domain.Order) *types.Order {
	return &types.Order{
		Model: gorm.Model{
			ID:        uint(orderDomain.ID),
			CreatedAt: orderDomain.CreatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(orderDomain.DeletedAt)),
		},
		UUID:          orderDomain.UUID.String(),
		SubmittedAt:   orderDomain.SubmittedAt,
		UserID:        uint(orderDomain.UserID),
		PaymentMethod: uint8(orderDomain.PaymentMethod),
	}
}

func OrderStorage2Domain(order types.Order) (*domain.Order, error) {
	uid, err := domain.OrderUUIDFromString(order.UUID)
	return &domain.Order{
		ID:            domain.OrderID(order.ID),
		UUID:          uid,
		CreatedAt:     order.CreatedAt,
		DeletedAt:     order.DeletedAt.Time,
		SubmittedAt:   order.SubmittedAt,
		UserID:        userDomain.UserID(order.UserID),
		PaymentMethod: domain.TypePaymentMethod(order.PaymentMethod),
	}, err
}
