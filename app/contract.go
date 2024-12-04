package app

import (
	"context"
	"order-service/config"
	orderPort "order-service/internal/order/port"
	userPort "order-service/internal/user/port"

	"gorm.io/gorm"
)

type App interface {
	OrderService(ctx context.Context) orderPort.Service
	UserService(ctx context.Context) userPort.Service
	DB() *gorm.DB
	Config() config.Config
}
