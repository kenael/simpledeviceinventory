package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"time"
)

func Server() {
	app := fiber.New()
	
	// Logging
	app.Use(logger.New(logger.Config{
		Format: "[${time}] from: ${ip} - ${status} - ${latency} ${method} ${url} ${query} ${header} ${path}\n",
		TimeFormat: time.RFC3339,
		TimeZone: "Europe/Berlin",
		Output: os.Stdout,
	}))

    // GET /api/register
    app.Get("/", helloWorld)

    app.Listen(":3000")
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World !")
}

