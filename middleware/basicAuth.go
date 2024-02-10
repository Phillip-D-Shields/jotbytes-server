package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// BasicAuth checks for a specific API key in the header
func BasicAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		apiKey := c.Get("123456")

		// Replace "your_api_key" with your actual API key value
		if apiKey != "123456" {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		return c.Next()
	}
}
