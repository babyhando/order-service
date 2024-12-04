package app

import (
	"context"

	"github.com/babyhando/order-service/config"
	orderPort "github.com/babyhando/order-service/internal/order/port"
	userPort "github.com/babyhando/order-service/internal/user/port"

	"gorm.io/gorm"
)

type App interface {
	OrderService(ctx context.Context) orderPort.Service
	UserService(ctx context.Context) userPort.Service
	DB() *gorm.DB
	Config() config.Config
}
