package routes

import (
	"hims/pkg/config"
	"hims/pkg/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func RegisterRoutes(r fiber.Router, cfg config.Config, db *bun.DB) {
	r.Get("/", handlers.TestAPI)
}
