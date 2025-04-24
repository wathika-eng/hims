package routes

import (
	"hims/pkg/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router) {
	r.Get("/", handlers.TestAPI)
}
