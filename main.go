package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/robertjbass/go-rest/endpoints"
)

func main() {
	app := fiber.New()

	app.Get("/", endpoints.RootHandler)

	app.Listen(":3000")
}
