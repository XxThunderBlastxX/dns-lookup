package routes

import (
	"github.com/XxThunderBlastxX/dns-lookup/controller"
	"github.com/XxThunderBlastxX/dns-lookup/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Router(app *fiber.App) {
	//Initial Route
	app.Get("/", middleware.RateLimiting(), controller.InitialRoute())

	//Fiber Monitor
	app.Get("/monitor", middleware.RateLimiting(), monitor.New(monitor.Config{Title: "DNS-Lookup"}))

	//name server route
	app.Get("/ns", middleware.RateLimiting(), controller.NsLookup())
	//cname route
	app.Get("/cname", middleware.RateLimiting(), controller.CnameLookup())
	//txt record route
	app.Get("/txt", middleware.RateLimiting(), controller.TxtLookup())
	//mx record route
	app.Get("/mx", middleware.RateLimiting(), controller.MxLookup())

}
