package controller

import "github.com/gofiber/fiber/v2"

func InitialRoute() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		jsonData := fiber.Map{"Name": "DNS - Lookup API", "Created By": "Koustav Mondal <ThunderBlast>", "Status": "Running", "Version": "0.0.1"}

		return ctx.Status(fiber.StatusOK).JSON(jsonData)
	}
}
