package utils

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

func GetPage(url string) (playwright.Page, error) {
	pw, err := playwright.Run()
	if err != nil {
		fmt.Printf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		fmt.Printf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		fmt.Printf("could not create page: %v", err)
	}
	if _, err = page.Goto(url); err != nil {
		fmt.Printf("could not goto: %v", err)
	}
	return page, err
}
