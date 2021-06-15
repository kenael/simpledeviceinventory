package server

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	//"encoding/json"
	"log"

	"fmt"

	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/go-gcfg/gcfg"
    	"database/sql"
	"github.com/kenael/simpledeviceinventory/env"
	"github.com/zcalusic/sysinfo"
)

type systemValues struct {
	MachineID string
	Packages  []env.Package
	SysInfo   sysinfo.SysInfo
}

func Server() {
	Cfg := struct {
		Server struct {
			Port string
		}
		Database struct {
			User string
			Password string
			DbServer string
			DbServerport string
			Database string
			Option string

		}
	}{}
	configFile := env.GetEnv("SIMPLEDEVINV_CONFIG_FILE", "server.conf")
	err := gcfg.FatalOnly(gcfg.ReadFileInto(&Cfg, configFile))
        if err != nil {
		log.Fatal("ERROR  can't read Configfile", err)
	}
	// start Connect DB
	DB, err = sql.Open("postgres", "postgres://" + Cfg.Database.User +
	  ":" + Cfg.Database.Password +
	  "@"+ Cfg.Database.DbServer + ":"+ Cfg.Database.DbServerport +
	  "/"+ Cfg.Database.Database + Cfg.Database.Option)
        if err != nil {
           log.Fatal(err)
        }

	log.Printf("Database Connect %s Port %s Database %s as User %s", Cfg.Database.DbServer, Cfg.Database.DbServerport, Cfg.Database.Database, Cfg.Database.User)
	// start Webserver
	app := fiber.New()
	app.Use(recover.New()) // Error handling
	// Logging
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] from: ${ip} - ${status} - ${latency} ${method} ${url} ${query} ${header} ${path}\n",
		TimeFormat: time.RFC3339,
		TimeZone:   "Europe/Berlin",
		Output:     os.Stdout,
	}))

	// GET /api/register
	app.Get("/", helloWorld)
	system := app.Group("/system")
	system.Post("/:machineID", receivSystem)
	system.Post("/:machineID/packages", receivPackages)
	system.Post("/:machineID/user", receivUser)
    app.Listen(":"+Cfg.Server.Port)
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
	StoreSystemValues(c.Params("machineID"), c.Body())
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
	StorePackagesValues(c.Params("machineID"), payload)
	// for key, item := range payload {
	// 	log.Printf("%v Item: %s Version: %s ",key, item.Package, item.Version)
	// }
	// log.Printf("Payload: %v ", payload)
	return c.SendString(fmt.Sprintf("Receive %v Pakackages Information", len(payload)))

}

func receivUser(c *fiber.Ctx) error {
	c.Accepts("application/json", "text")
	log.Printf("Param machineID: %s\n", c.Params("machineID"))

	StoreSystemUser(c.Params("machineID"), c.Body())
	
	return c.SendString("ok")
}
