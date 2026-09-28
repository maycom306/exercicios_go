package router

import (
	"ex07/middleware"

	"github.com/gofiber/fiber/v3"
)

func RouterGO(app *fiber.App) {
	app.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"Status": "ok",
		})
	})

	app.Get("/info", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"app":    "ex07",
			"versao": "1.0",
		})
	})

	admin := app.Group("/admin", middleware.RequireAPIKey())
	usuarios := []Usuarios{
		{ID: 1, Nome: "Michael"},
		{ID: 2, Nome: "Willieny"},
	}
	
	proximoID := 3

	admin.Get("/usuarios", func(c fiber.Ctx) error {
		return c.JSON(usuarios)

	})
	admin.Post("/usuarios", func(c fiber.Ctx) error {
		u := new(Usuarios)
		if err := c.Bind().Body(u); err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"erro": "JSON inválido",
			})
		}

		u.ID = proximoID
		proximoID++
		usuarios = append(usuarios, *u)

		return c.Status(fiber.StatusCreated).JSON(u)
	})
}
