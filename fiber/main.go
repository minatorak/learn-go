package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(csrf.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path} ${body} -> ${resBody} \n",
	}))

	app.Get("/:value", func(c *fiber.Ctx) error {
		result := c.Params("value")

		return c.SendString("Hello, World!" + result)
	})

	app.Listen(":3000")

}
