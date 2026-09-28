package main

import (
	"log"

	"ex07/middleware"
	"ex07/router"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("erro ao carregar .env: ", err)
	}

	app := fiber.New()
	app.Use(middleware.LoggerTime())
	app.Use(middleware.Cors())

	router.RouterGO(app)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
