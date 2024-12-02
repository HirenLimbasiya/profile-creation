package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LatencyLogger(c *fiber.Ctx)error{
	start := time.Now()
	err := c.Next()
	log.Printf("Latency: %v", time.Since(start))
	return err
}