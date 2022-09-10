package utils

import (
	"github.com/playwright-community/playwright-go"
)

func GetValueFromPage(url string, className string) (any, error) {

	/*
		Run the playwright
	*/
	pw, err := playwright.Run()
	if err != nil {
		return nil, err
	}

	/*
		Prepare the browser and create page
	*/
	browser, err := pw.Chromium.Launch()
	if err != nil {
		return nil, err
	}
	page, err := browser.NewPage()
	if err != nil {
		return nil, err
	}

	/*
		Go to the target page and get the value
	*/
	if _, err = page.Goto(url); err != nil {
		return nil, err
	}
	entries, err := page.QuerySelectorAll(className)
	if err != nil {
		return nil, err
	}
	title, err := entries[0].InnerHTML()
	if err != nil {
		return nil, err
	}

	/*
	 Close the playwright
	*/
	if err = browser.Close(); err != nil {
		return nil, err
	}
	if err = pw.Stop(); err != nil {
		return nil, err
	}
	return title, nil
}
