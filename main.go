package main

import (
	"fmt"
	"os"
	"profile-creation/api"
	"profile-creation/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	err := godotenv.Load()
	if err!= nil {
        fmt.Println("Error loading.env file")
    }
	port := os.Getenv("PORT");
	app.Use(middleware.LatencyLogger)
	api.SetupRoutes(app)

	if port == "" {
        port = "3000"
    }
	fmt.Println(port)
	app.Listen(":" + port)
}