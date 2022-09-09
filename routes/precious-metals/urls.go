package preciousmetals

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {
	api := app.Group("/precious-metal")
	api.Get("/", PreciousMetals)
}
