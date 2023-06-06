package main

import (
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/golang-jwt/jwt/v5"
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

	app.Get("/hello/:value", func(c *fiber.Ctx) error {
		result := c.Params("value")

		return c.SendString("Hello, World!" + result)
	})
	// Login route
	app.Get("/login", login)

	// Unauthenticated route
	app.Get("/", accessible)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret!!!!!")},
	}))

	// Restricted Routes
	app.Get("/restricted", restricted)

	app.Listen(":3000")

}

func login(c *fiber.Ctx) error {
	// user := c.FormValue("user")
	// pass := c.FormValue("pass")

	// // Throws Unauthorized error
	// if user != "john" || pass != "doe" {
	// 	return c.SendStatus(fiber.StatusUnauthorized)
	// }

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Second * 60).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret!!!!!"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}
