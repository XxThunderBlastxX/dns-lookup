package controller

import (
	"github.com/XxThunderBlastxX/dns-lookup/models"
	"github.com/gofiber/fiber/v2"
	"net"
)

// NsLookup is a handler to get and show the name servers of a domain
func NsLookup() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var domain models.Request
		var nsResponse models.ResponseNs

		err := ctx.BodyParser(&domain)
		if err != nil {
			parseErr := models.ResponseErr{}
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(parseErr.PresentError(err))
		}

		nameServers, _ := net.LookupNS(domain.Domain)
		for _, ns := range nameServers {
			nsResponse.NameServer = append(nsResponse.NameServer, ns.Host)
		}

		return ctx.Status(fiber.StatusOK).JSON(nsResponse.PresentNs(nsResponse))
	}
}
