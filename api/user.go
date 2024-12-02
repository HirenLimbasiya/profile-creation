package api

import (
	"profile-creation/types"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error{
	var user types.UserRequest

	if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload", "data": nil})
    }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User created successfully", "data": user})
}