package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Dummy handler
func handler() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendString("This is a dummy route")
	}
}

//// Dummy middleware
//func middleware() func(*fiber.Ctx) error {
//return func(c *fiber.Ctx) error {
//return c.Next()
//}
//}

func main() {
	app := fiber.New()

	// Serve static files from the ./public directory
	app.Static("/", "./public")

	// Basic handler
	app.Get("/greeting", func(c *fiber.Ctx) error {
		if _, err := c.WriteString("Hello, World!"); err != nil { // => "Hello, World!"
			return err
		}

		// WriteString adopts the string
		_, err := fmt.Fprintf(c, "%s\n", "Hello, World!") // => "Hello, World!Hello, World!"
		return err
	})

	// Basic handler with params
	app.Get("/greeting/:value", func(c *fiber.Ctx) error {
		return c.SendString("Hello, " + c.Params("value") + "!")
	})

	// Grouping
	api := app.Group("/api", logger.New( // Middleware to log useful data
		logger.Config{
			Format: "http://localhost:${port}${path}\n", // NOTE: Port behaves weird in logs
		},
	))

	v1 := api.Group("/v1")     // /api/v1
	v1.Get("/list", handler()) // /api/v1/list
	v1.Get("/user", handler()) // /api/v1/user

	v2 := api.Group("/v2")     // /api/v2
	v2.Get("/list", handler()) // /api/v2/list
	v2.Get("/user", handler()) // /api/v2/user

	app.Use(func(c *fiber.Ctx) error { // Middleware to match anything
		return c.SendStatus(404) // => 404 "Not Found"
	})

	app.Listen(":3000")
}
