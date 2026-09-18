package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
)

func RequireApi() fiber.Handler {
	return func(c fiber.Ctx) error {
		key := c.Get("X_API_key")

		if key != apiKey {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Chave de API Invalida",
			})

		}
		return c.Next()
	}
}

func LoggerTime() fiber.Handler {
	return func(c fiber.Ctx) error {
		horario_inicial := time.Now()

		erro := c.Next()

		duracao := time.Since(horario_inicial)
		fmt.Printf("O metodo é [%s].\n Caminho: /%s\n Status: %d\n A duração total foi: %v", c.Method(), c.Path(), c.Response().StatusCode(), duracao)

		return erro
	}

}

func Cors() fiber.Handler {
	return func(c fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin","*")
		if c.Method() == "OPTIONS"{
			return c.SendStatus(200)
		}
		return c.Next()
	}
}
