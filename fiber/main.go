package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Get("/:value", func(c *fiber.Ctx) error {
		result := c.Params("value")

		return c.SendString("Hello, World!" + result)
	})

	app.Listen(":3000")

}
