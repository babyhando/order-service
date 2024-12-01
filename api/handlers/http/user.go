package http

import (
	"errors"
	"order-service/api/pb"
	"order-service/api/service"

	"github.com/gofiber/fiber/v2"
)

func SignUp(svc *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req pb.UserSignUpRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		resp, err := svc.SignUp(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, service.ErrUserCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

func GetUserByID(svc *service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		resp, err := svc.GetByID(c.UserContext(), uint(id))
		if err != nil {
			if errors.Is(err, service.ErrUserNotFound) {
				return c.SendStatus(fiber.StatusNotFound)
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}
