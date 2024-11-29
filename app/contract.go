package app

import (
	"order-service/config"
	orderPort "order-service/internal/order/port"
	userPort "order-service/internal/user/port"
)

type App interface {
	OrderService() orderPort.Service
	UserService() userPort.Service
	Config() config.Config
}
