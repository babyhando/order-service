package app

import (
	"fmt"
	"order-service/config"
	"order-service/internal/order"
	orderPort "order-service/internal/order/port"
	"order-service/internal/user"
	userPort "order-service/internal/user/port"
	"order-service/pkg/adapters/storage"
	"order-service/pkg/cache"
	"order-service/pkg/postgres"

	redisAdapter "order-service/pkg/adapters/cache"

	"gorm.io/gorm"
)

type app struct {
	db            *gorm.DB
	cfg           config.Config
	orderService  orderPort.Service
	userService   userPort.Service
	redisProvider cache.Provider
}

func (a *app) OrderService() orderPort.Service {
	return a.orderService
}

func (a *app) UserService() userPort.Service {
	return a.userService
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) setDB() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   a.cfg.DB.User,
		Pass:   a.cfg.DB.Password,
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		DBName: a.cfg.DB.Database,
		Schema: a.cfg.DB.Schema,
	})

	if err != nil {
		return err
	}

	a.db = db
	return nil
}

func (a *app) setRedis() {
	a.redisProvider = redisAdapter.NewRedisProvider(fmt.Sprintf("%s:%d", a.cfg.Redis.Host, a.cfg.Redis.Port))
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{
		cfg: cfg,
	}

	if err := a.setDB(); err != nil {
		return nil, err
	}

	a.setRedis()

	a.orderService = order.NewService(nil, storage.NewOrderRepo(a.db))
	a.userService = user.NewService(storage.NewUserRepo(a.db, true, a.redisProvider))

	return a, nil
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
