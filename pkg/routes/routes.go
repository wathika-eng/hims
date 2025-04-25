package routes

import (
	"hims/pkg/config"
	"hims/pkg/handlers"
	"hims/pkg/middlewares"
	"hims/pkg/repo"
	"hims/pkg/services"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func RegisterRoutes(r fiber.Router, cfg *config.Config, db *bun.DB) {
	repo := repo.NewRepo(db)
	// create tables panics if err, so as to stop app
	err := repo.Up()
	if err != nil {
		panic(err)
	}
	// err := repo.Reset()
	// if err != nil {
	// 	panic(err)
	// }

	services := services.NewServices(repo)
	handlers := handlers.NewHandler(repo, services)

	r.Get("/", handlers.TestAPI)
	r.Post("api/signup", handlers.NewDoctor)
	r.Post("api/login", handlers.LoginDoctor)
	doc := r.Group("/api/")
	doc.Use(middlewares.Auth)
	{

		doc.Post("/add-patient", handlers.NewPatient)
	}
}
