package endpoints

import (
	"github.com/gofiber/fiber/v2"
)

func RootHandler(c *fiber.Ctx) error {
	return c.SendString("Root Handler")
}
