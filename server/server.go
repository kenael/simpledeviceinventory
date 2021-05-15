package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"time"
	//"encoding/json"
	"github.com/kenael/simpledeviceinventory/env"
	"log"
	"github.com/zcalusic/sysinfo"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type systemValues struct {
	MachineID string 
	Packages []env.Package 
	SysInfo sysinfo.SysInfo
}

func Server() {
	port := env.GetEnv("SERVER_PORT", "3000")

	app := fiber.New()
	app.Use(recover.New()) // Error handling
	// Logging
	app.Use(logger.New(logger.Config{
		Format: "[${time}] from: ${ip} - ${status} - ${latency} ${method} ${url} ${query} ${header} ${path}\n",
		TimeFormat: time.RFC3339,
		TimeZone: "Europe/Berlin",
		Output: os.Stdout,
	}))

    // GET /api/register
    app.Get("/", helloWorld)
	system := app.Group("/system")
	system.Post("/:machineID", receivSystem)
	system.Post("/:machineID/packages", receivPackages)
    app.Listen(":"+port)
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World !")
}

func receivSystem(c *fiber.Ctx) error {
	c.Accepts("application/json", "text")
	log.Printf("Param machineID: %s\n", c.Params("machineID"))
	var payload sysinfo.SysInfo


	err := c.BodyParser(&payload)
	if err != nil {
		log.Println("Parser Error")
		return fiber.NewError(fiber.StatusInternalServerError, "Parser Error")
	}
	log.Printf("Hostname: %s  ", payload.Node.Hostname)
	return c.JSON(payload)
}

func receivPackages(c *fiber.Ctx) error {
	c.Accepts("application/json", "text")
	log.Printf("Param machineID: %s\n", c.Params("machineID"))

	var payload []env.Package
	err := c.BodyParser(&payload)
	if err != nil {
		log.Println("Parser Error")
		
	}
	for key, item := range payload {
		log.Printf("%v Item: %s Version: %s ",key, item.Package, item.Version)
	}
	// log.Printf("Payload: %v ", payload)
	return c.SendString("ok")


}