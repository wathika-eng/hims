package api

import (
	"context"
	"fmt"
	"hims/pkg/config"
	"hims/pkg/database"
	"hims/pkg/routes"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

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
	// close db once we exit
	defer db.Close()

	// initialize a new fiber app instance
	app := fiber.New(fiber.Config{
		AppName: "HISM V1.0.0",
	})
	app.Use(logger.New())
	app.Use(recover.New())
	// app.Use(helmet.New())
	routes.RegisterRoutes(app, cfg, db)
	if err := app.Listen(fmt.Sprintf(":%v", cfg.PORT)); err != nil {
		log.Fatalf("error while starting the server: %v", err)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		if err := app.Listen(fmt.Sprintf(":%v", cfg.PORT)); err != nil {
			log.Fatalf("error while starting the server: %v", err)
		}
	}()
	<-ctx.Done()
	log.Println("shutting down...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("failed to shutdown fiber: %v", err)
	}
}
