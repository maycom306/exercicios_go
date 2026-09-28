package middleware

import (
	"os"
	"github.com/gofiber/fiber/v3"
)

func RequireAPIKey() fiber.Handler {
	return func(c fiber.Ctx) error {
		key := c.Get("X-API-Key")
		var apiKey = os.Getenv("API_KEY")

		if key != apiKey {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Chave de API Invalida",
			})

		}
		return c.Next()
	}
}
