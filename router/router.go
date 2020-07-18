package router

import (
	"github.com/gofiber/fiber"
	"github.com/komfluent/imgen/handler"
)

// Setup api routes
func RegisterRoutes(app *fiber.App) {
	// Setup index
	app.Get("/", handler.Index)
	// Setup api group
	api := app.Group("/api")
	// Version 1
	v1 := api.Group("/v1")

	// Setup Routes
	v1.Get("/", func(c *fiber.Ctx) { c.Status(501).Send("Unimplemented") })

}
