package mapper

import (
	"order-service/internal/order/domain"
	userDomain "order-service/internal/user/domain"
	"order-service/pkg/adapters/storage/types"
)

func OrderDomain2Storage(orderDomain domain.Order) *types.Order {
	return &types.Order{
		ID:            uint(orderDomain.ID),
		UUID:          orderDomain.UUID.String(),
		CreatedAt:     orderDomain.CreatedAt,
		DeletedAt:     orderDomain.DeletedAt,
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
		DeletedAt:     order.DeletedAt,
		SubmittedAt:   order.SubmittedAt,
		UserID:        userDomain.UserID(order.UserID),
		PaymentMethod: domain.TypePaymentMethod(order.PaymentMethod),
	}, err
}
