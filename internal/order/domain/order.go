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

type Order struct {
	ID            OrderID
	UUID          string
	CreatedAt     time.Time
	DeletedAt     time.Time
	SubmittedAt   time.Time
	UserID        userDomain.UserID
	PaymentMethod TypePaymentMethod
	Items         []OrderItem
}

type OrderItemID uint

type OrderItem struct {
	ID OrderItemID
	//
}

type OrderListFilters struct {
	//
}
