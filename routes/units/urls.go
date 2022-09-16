package units

import "github.com/gofiber/fiber/v2"

func Setup(app fiber.Router) {
	api := app.Group("/units")
	api.Get("/", MoneyUnit)
}
