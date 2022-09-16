package preciousmetals

import (
	"currencyapi/utils"

	"github.com/gofiber/fiber/v2"
)

func PreciousMetals(c *fiber.Ctx) error {
	// title, err := utils.GetValueFromPage(GoldExchangeUrl, "input.w-button")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return c.Status(400).JSON(utils.JsonResponse{Msg: err.Error(), Result: []any{}})
	// }
	return c.Status(200).JSON(
		utils.JsonResponse{
			Msg:    "ok",
			Result: "",
		})
}
