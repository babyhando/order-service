package domain

import (
	userDomain "order-service/internal/user/domain"
	"time"

	"github.com/google/uuid"
)

type TypePaymentMethod uint8

const (
	PaymentMethodUnknown TypePaymentMethod = iota
	PaymentTypeIPG
	PaymentTypeWallet
	PaymentTypeMix
)

type (
	OrderID   uint
	OrderUUID = uuid.UUID
)

func OrderUUIDFromString(s string) (OrderUUID, error) {
	uid, err := uuid.Parse(s)
	return OrderUUID(uid), err
}

type Order struct {
	ID            OrderID
	UUID          OrderUUID
	CreatedAt     time.Time
	DeletedAt     time.Time
	SubmittedAt   time.Time
	UserID        userDomain.UserID
	PaymentMethod TypePaymentMethod
	Items         []OrderItem
}

type OrderItemID uint

type OrderItem struct {
	ID          OrderItemID
	ProductName string
	UnitPrice   uint
	Quantity    uint
	Description string
}

type OrderListFilters struct {
	//
}
