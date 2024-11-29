package types

import "time"

type Order struct {
	ID            uint
	UUID          string
	CreatedAt     time.Time
	DeletedAt     time.Time
	UpdatedAt     time.Time
	SubmittedAt   time.Time
	UserID        uint
	PaymentMethod uint8
}

type OrderItem struct {
	ID          uint
	ProductName string
	UnitPrice   uint
	Quantity    uint
	Description string
}
