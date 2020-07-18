package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/komfluent/imgen/router"
)

func main() {
	app := fiber.New()

	router.RegisterRoutes(app)

	log.Fatal(app.Listen(3000))
}
