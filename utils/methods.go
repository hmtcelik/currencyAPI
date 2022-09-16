package utils

import "github.com/gofiber/fiber/v2"

func ReturnErrorIfError(err error, status int, c *fiber.Ctx) error {
	if err != nil {
		return c.Status(status).JSON(
			JsonResponse{
				Msg:    err.Error(),
				Result: "",
			})
	} else {
		return nil
	}
}

// func GetValueFromPage(url string, className string) (any, error) {

// 	/*
// 		Run the playwright
// 	*/
// 	pw, err := playwright.Run()
// 	if err != nil {
// 		return nil, err
// 	}

// 	/*
// 		Prepare the browser and create page
// 	*/
// 	browser, err := pw.Chromium.Launch()
// 	if err != nil {
// 		return nil, err
// 	}
// 	page, err := browser.NewPage()
// 	if err != nil {
// 		return nil, err
// 	}

// 	/*
// 		Go to the target page and get the value
// 	*/
// 	if _, err = page.Goto(url); err != nil {
// 		return nil, err
// 	}
// 	entries, err := page.QuerySelectorAll(className)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var title string = ""
// 	if len(entries) != 0 {
// 		title, err = entries[0].InnerHTML()
// 		if err != nil {
// 			return nil, err
// 		}
// 	} else {
// 		return nil, errors.New("Did not find")
// 	}
// 	/*
// 	 Close the playwright
// 	*/
// 	if err = browser.Close(); err != nil {
// 		return nil, err
// 	}
// 	if err = pw.Stop(); err != nil {
// 		return nil, err
// 	}
// 	return title, nil
// }
