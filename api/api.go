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
	"time"

	"github.com/gofiber/fiber/v3/middleware/static"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
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
	app := fiber.New(fiber.Config{AppName: "HISM V1.0.0"})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(helmet.New(helmet.Config{
		XSSProtection:         "1",
		ContentSecurityPolicy: "",
	}))
	app.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173/"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	app.Get("/*", static.New("./frontend/dist", static.Config{
		//Compress:      true,
		ByteRange:     true,
		Browse:        true,
		IndexNames:    []string{"index.html"},
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	}))

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
