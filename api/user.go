package api

import (
	"profile-creation/types"
	"profile-creation/validation"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error{
	var user types.UserRequest

	if err := c.BodyParser(&user); err != nil {
       return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"error":       "Invalid request payload",
			"data":        nil,
		})
    }

	if err := validation.ValidateUserInput(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status_code": fiber.StatusBadRequest,
			"error":       err.Error(),
			"data":        nil,
		})
    }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status_code": fiber.StatusOK,
		"message":     "User created successfully",
		"data":        user,
	})
}