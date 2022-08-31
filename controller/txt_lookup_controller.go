package controller

import (
	"github.com/XxThunderBlastxX/dns-lookup/models"
	"github.com/gofiber/fiber/v2"
	"net"
)

// TxtLookup is a handler to fetch all the txt record for a domain
func TxtLookup() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var domain models.Request
		var txtResponse models.ResponseTxt

		if err := ctx.BodyParser(&domain); err != nil {
			reqErr := models.ResponseErr{Error: err.Error()}
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(reqErr)
		}

		txtRecords, _ := net.LookupTXT("facebook.com")
		for _, txt := range txtRecords {
			txtResponse.Txt = append(txtResponse.Txt, txt)
		}

		return ctx.Status(fiber.StatusOK).JSON(txtResponse)
	}
}
