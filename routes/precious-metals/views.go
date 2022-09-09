package preciousmetals

import (
	"currencyapi/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func PreciousMetals(c *fiber.Ctx) error {
	page, err := utils.GetPage("https://abdulhamitcelik.com/")
	if err != nil {
		return c.Status(400).JSON(utils.JsonResponse{Msg: "Error", Result: []any{}})
	}
	titleVal, err := page.QuerySelectorAll(".title")
	if err != nil {
		fmt.Printf("could not get entries: %v", err)
		return c.Status(400).JSON(utils.JsonResponse{Msg: "Error", Result: []any{}})
	}
	return c.Status(200).JSON(
		utils.JsonResponse{
			Msg:    "ok",
			Result: titleVal,
		})
}
