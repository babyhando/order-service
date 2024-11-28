package app

import (
	"order-service/config"
	orderPort "order-service/internal/order/port"
)

type App interface {
	OrderService() orderPort.Service
	Config() config.Config
}
