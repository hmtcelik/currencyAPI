package preciousmetals

import (
	"currencyapi/utils"

	"github.com/gofiber/fiber/v2"
)

func PreciousMetals(c *fiber.Ctx) error {
	return c.Status(200).JSON(
		utils.JsonResponse{
			Msg: "ok",
		})
}
