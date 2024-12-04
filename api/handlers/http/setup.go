package http

import (
	"fmt"
	"order-service/app"
	"order-service/config"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	router := fiber.New()

	api := router.Group("/api/v1")

	registerAuthAPI(appContainer, cfg, api)

	return router.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}

func registerAuthAPI(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	userSvcGetter := userServiceGetter(appContainer, cfg)
	router.Post("/sign-up", SignUp(userSvcGetter), setTransaction(appContainer.DB()))
	router.Post("/sign-in", SignIn(userSvcGetter), setTransaction(appContainer.DB()))
}
