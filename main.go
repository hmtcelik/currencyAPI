package main

import (
	preciousmetals "currencyapi/routes/precious-metals"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/playwright-community/playwright-go"
)

func main() {
	app := fiber.New()

	err := playwright.Install()
	if err != nil {
		fmt.Println("The playwright didnt install correctly")
		return
	}

	preciousmetals.Setup(app)

	app.Listen(":3000")
}
