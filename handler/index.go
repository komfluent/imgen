package handler

import "github.com/gofiber/fiber"

// Index page handler
func Index(c *fiber.Ctx) {
	c.Send("Hello World")
}
