package api

import (
	"hims/pkg/config"
	"hims/pkg/database"
	"hims/pkg/routes"
	"log"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewServer() {
	// default logger
	logg := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logg)
	// load env variables before app starts
	cfg, err := config.LoadConfigs()
	// if it returns an error, exit early
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// initialize a new fiber app instance
	app := fiber.New(fiber.Config{
		ServerHeader: "hism",
		AppName:      "HISM V1.0.0",
	})
	app.Use(logger.New())
	app.Use(recover.New())
	//r.Use(helmet.New())c
	routes.RegisterRoutes(app, cfg, db)

	if err := app.Listen(":8000"); err != nil {
		log.Fatalf("error while starting the server: %v", err)
	}
}
