package routes

import (
	"profile-creation/api"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/user", api.CreateUser)
}