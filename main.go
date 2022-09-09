package main

import (
	preciousmetals "currencyapi/routes/precious-metals"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	preciousmetals.Setup(app)

	app.Listen(":3000")
}
