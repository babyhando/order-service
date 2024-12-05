package http

import (
	"context"

	"github.com/babyhando/order-service/api/service"
	"github.com/babyhando/order-service/app"
	"github.com/babyhando/order-service/config"
)

// user service transient instance handler
func userServiceGetter(appContainer app.App, cfg config.ServerConfig) ServiceGetter[*service.UserService] {
	return func(ctx context.Context) *service.UserService {
		return service.NewUserService(appContainer.UserService(ctx),
			cfg.Secret, cfg.AuthExpMinute, cfg.AuthRefreshMinute)
	}
}
