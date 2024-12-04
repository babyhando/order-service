package http

import (
	"context"
	"order-service/api/service"
	"order-service/app"
	"order-service/config"
)

func userServiceGetter(appContainer app.App, cfg config.ServerConfig) ServiceGetter[*service.UserService] {
	return func(ctx context.Context) *service.UserService {
		return service.NewUserService(appContainer.UserService(ctx),
			cfg.Secret, cfg.AuthExpMinute, cfg.AuthRefreshMinute)
	}
}
