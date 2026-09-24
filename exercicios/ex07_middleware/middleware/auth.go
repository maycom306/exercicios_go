package middleware

import (
	"os"
	"github.com/gofiber/fiber/v3"
)

var apiKey = os.Getenv("API_KEY")

func RequireApi() fiber.Handler {
	return func(c fiber.Ctx) error {
		key := c.Get("X-API-Key")

		if key != apiKey {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Chave de API Invalida",
			})

		}
		return c.Next()
	}
}



