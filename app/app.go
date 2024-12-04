package app

import (
	"context"
	"fmt"

	"github.com/babyhando/order-service/config"
	"github.com/babyhando/order-service/internal/order"
	orderPort "github.com/babyhando/order-service/internal/order/port"
	"github.com/babyhando/order-service/internal/user"
	userPort "github.com/babyhando/order-service/internal/user/port"
	"github.com/babyhando/order-service/pkg/adapters/storage"
	"github.com/babyhando/order-service/pkg/cache"
	"github.com/babyhando/order-service/pkg/postgres"

	redisAdapter "github.com/babyhando/order-service/pkg/adapters/cache"

	"gorm.io/gorm"

	appCtx "github.com/babyhando/order-service/pkg/context"
)

type app struct {
	db            *gorm.DB
	cfg           config.Config
	orderService  orderPort.Service
	userService   userPort.Service
	redisProvider cache.Provider
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) OrderService(ctx context.Context) orderPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.orderService == nil {
			a.orderService = a.orderServiceWithDB(a.db)
		}
		return a.orderService
	}

	return a.orderServiceWithDB(db)
}

func (a *app) orderServiceWithDB(db *gorm.DB) orderPort.Service {
	return order.NewService(a.userServiceWithDB(db), storage.NewOrderRepo(db))
}

func (a *app) UserService(ctx context.Context) userPort.Service {
	db := appCtx.GetDB(ctx)
	if db == nil {
		if a.userService == nil {
			a.userService = a.userServiceWithDB(a.db)
		}
		return a.userService
	}

	return a.userServiceWithDB(db)
}

func (a *app) userServiceWithDB(db *gorm.DB) userPort.Service {
	return user.NewService(storage.NewUserRepo(db, true, a.redisProvider))
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

	return a, nil
}

func NewMustApp(cfg config.Config) App {
	app, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return app
}
