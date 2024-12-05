package http

import (
	"fmt"

	"github.com/babyhando/order-service/app"
	"github.com/babyhando/order-service/config"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	router := fiber.New()

	api := router.Group("/api/v1", setUserContext)

	registerAuthAPI(appContainer, cfg, api)

	return router.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}

func registerAuthAPI(appContainer app.App, cfg config.ServerConfig, router fiber.Router) {
	userSvcGetter := userServiceGetter(appContainer, cfg)
	router.Post("/sign-up", setTransaction(appContainer.DB()), SignUp(userSvcGetter))
	router.Post("/sign-in", setTransaction(appContainer.DB()), SignIn(userSvcGetter))
}
