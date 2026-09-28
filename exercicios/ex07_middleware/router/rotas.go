package router

import "github.com/gofiber/fiber/v3"

func RouterGO(app *fiber.App) {
	app.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"Status": "ok",
		})
	})

	app.Get("/info", func (c fiber.Ctx)error{
		return c.JSON(fiber.Map{
			"app": "ex07",
			"versao": "1.0",
		})
	})

}
func 
