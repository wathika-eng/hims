package routes

import (
	"hims/pkg/config"
	"hims/pkg/handlers"
	"hims/pkg/repo"
	"hims/pkg/services"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func RegisterRoutes(r fiber.Router, cfg *config.Config, db *bun.DB) {
	repo := repo.NewRepo(db)
	services := services.NewServices(repo)
	handlers := handlers.NewHandler(repo, services)

	r.Get("/", handlers.TestAPI)
}
