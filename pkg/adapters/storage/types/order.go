package types

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UUID          string
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
