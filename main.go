package main

import (
	"currencyapi/routes/units"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")
	units.Setup(api)

	app.Listen(":8888")
}
