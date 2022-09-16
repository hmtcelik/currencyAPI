package main

import (
	preciousmetals "currencyapi/routes/precious-metals"
	"currencyapi/routes/units"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")
	preciousmetals.Setup(api)
	units.Setup(api)

	app.Listen(":3000")
}
