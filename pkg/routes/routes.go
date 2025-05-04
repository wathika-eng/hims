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
	// panic will be useful to run deferred functions while exiting

	err := repo.Up()
	if err != nil {
		panic(err)
	}
	// err := repo.Reset()
	// if err != nil {
	// 	panic(err)
	// }

	services := services.NewServices(repo, cfg)
	handler := handlers.NewHandler(repo, services)

	api := r.Group("/api/v1")
	api.Get("/health", handler.TestAPI)
	api.Post("/auth/signup", handler.NewDoctor)
	api.Post("/auth/login", handler.LoginDoctor)

	protected := api.Group("/protected")
	protected.Use(middlewares.Auth)

	protected.Post("/programs", handler.AddProgram)
	protected.Get("/programs", handler.GetPrograms)
	protected.Get("/programs/pdf", handler.GeneratePDF)

	protected.Post("/patients", handler.NewPatient)
	protected.Get("/patients", handler.GetPatients)
	protected.Get("/patients/:id<[0-9]+>", handler.Profile)
	protected.Post("patients/enroll", handler.AddPatientProgram)
}
