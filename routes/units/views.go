package units

import (
	"currencyapi/routes/units/types/models"
	"currencyapi/routes/units/types/outputs"
	"currencyapi/routes/units/validator"
	"currencyapi/utils"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MoneyUnitQueryParams struct {
	Type string `query:"type,required"`
}

func MoneyUnit(c *fiber.Ctx) error {
	/*
		Get params and parse them
	*/
	params := new(MoneyUnitQueryParams)
	err := c.QueryParser(params)
	utils.ReturnErrorIfError(err, 400, c)

	/*
		Validate Params
	*/
	if !validator.IsValidUnit(params.Type) {
		utils.ReturnErrorIfError(errors.New(fmt.Sprintf("%s is Not Found", params.Type)), 404, c)
	}

	/*
		Find Uri and request
	*/
	var uri string = TCMB_URI
	year, month, day := time.Now().AddDate(0, 0, -1).Date()
	res, err := http.Get(fmt.Sprintf(uri, year, int(month), day, int(month), year))
	utils.ReturnErrorIfError(err, 400, c)

	/*
		Response Body To Byte
	*/
	defer res.Body.Close()
	var data models.TcmbVeri
	body, err := io.ReadAll(res.Body)
	utils.ReturnErrorIfError(err, 400, c)

	/*
		Xml Unmarshal
	*/
	err = xml.Unmarshal(body, &data)
	utils.ReturnErrorIfError(err, 400, c)

	/*
		Read Values And Return Final Data
	*/
	var finalData outputs.MoneyUnitOutput
	for _, item := range data.DovizKurListe.Kur {
		if item.DovizCinsi == params.Type {
			finalData.Type = item.DovizCinsi
			finalData.Value = item.Alis
		}
	}

	return c.Status(200).JSON(
		utils.JsonResponse{
			Msg:    "ok",
			Result: [...]outputs.MoneyUnitOutput{finalData},
		})
}
