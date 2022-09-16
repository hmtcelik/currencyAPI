package preciousmetals

import "github.com/gofiber/fiber/v2"

func Setup(app fiber.Router) {
	api := app.Group("/precious-metal")
	api.Get("/", PreciousMetals)
}
