package http

import (
	"fmt"
	"order-service/api/service"
	"order-service/app"
	"order-service/config"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	router := fiber.New()

	api := router.Group("/api/v1")

	userService := service.NewUserService(appContainer.UserService(),
		cfg.Secret, cfg.AuthExpMinute, cfg.AuthRefreshMinute)

	api.Post("/sign-up", SignUp(userService))

	api.Post("/sign-in", SignIn(userService))

	return router.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}
