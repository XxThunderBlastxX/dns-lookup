package controller

import (
	"github.com/XxThunderBlastxX/dns-lookup/models"
	"github.com/gofiber/fiber/v2"
	"net"
)

// CnameLookup is a handler to get and show the CNAME records of a domain
func CnameLookup() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var domain models.Request

		err := ctx.BodyParser(&domain)
		if err != nil {
			parseErr := models.ResponseErr{Error: err.Error()}
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(parseErr.PresentError())
		}

		cname, _ := net.LookupCNAME(domain.Domain)

		cnamePres := models.ResponseCname{Cname: cname}

		return ctx.Status(fiber.StatusOK).JSON(cnamePres.PresentCname())
	}
}
