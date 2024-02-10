package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// RequestLogger logs the details of the request
func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next() // Proceed to next middleware or handler
		duration := time.Since(start)

		println("Request:", c.Method(), c.Path(), "took", duration, "error:", err)
		return nil
	}
}
