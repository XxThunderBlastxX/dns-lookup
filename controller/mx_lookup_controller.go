package controller

import (
	"github.com/XxThunderBlastxX/dns-lookup/models"
	"github.com/gofiber/fiber/v2"
	"net"
)

// MxLookup is a handler to fetch all the mx record for a domain
func MxLookup() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var domain models.Request
		var mxResponse models.ResponseMx

		if err := ctx.BodyParser(&domain); err != nil {
			reqErr := models.ResponseErr{Error: err.Error()}
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(reqErr)
		}

		mxRecords, _ := net.LookupMX(domain.Domain)
		for _, mx := range mxRecords {
			mxResponse.Mx = append(mxResponse.Mx, mx.Host)
		}

		return ctx.Status(fiber.StatusOK).JSON(mxResponse.PresentMx())
	}
}
