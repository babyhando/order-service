package app

import (
	"order-service/config"
	"order-service/internal/order"
	orderPort "order-service/internal/order/port"
	"order-service/pkg/addapters/storage"

	"gorm.io/gorm"
)

type app struct {
	db           *gorm.DB
	cfg          config.Config
	orderService orderPort.Service
}

func (a *app) OrderService() orderPort.Service {
	return a.orderService
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) setDB() error {
	// sets a.sb
	return nil
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	a.orderService = order.NewService(nil, storage.NewOrderRepo(a.db))

	return a, nil
}
