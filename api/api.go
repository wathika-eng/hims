package api

import (
	"hims/pkg/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewServer() {
	// initialize a new fiber app instance
	app := fiber.New(fiber.Config{
		ServerHeader: "hism",
		AppName:      "HISM V1.0.0",
	})
	app.Use(logger.New())
	app.Use(recover.New())
	//r.Use(helmet.New())
	routes.RegisterRoutes(app)

	if err := app.Listen(":8000"); err != nil {
		log.Fatalf("error while starting the server: %v", err)
	}
}
