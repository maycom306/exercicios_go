package middleware

import (
	"fmt"
	"time"
	"github.com/gofiber/fiber/v3"
)

func LoggerTime() fiber.Handler {
	return func(c fiber.Ctx) error {
		horario_inicial := time.Now()

		erro := c.Next()

		duracao := time.Since(horario_inicial)
		fmt.Printf("O metodo é [%s].\n Caminho: /%s\n Status: %d\n A duração total foi: %v", c.Method(), c.Path(), c.Response().StatusCode(), duracao)

		return erro
	}
}