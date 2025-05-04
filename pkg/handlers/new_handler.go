// package handlers connects with the routes
package handlers

import (
	"hims/pkg/repo"
	"hims/pkg/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	repo     *repo.Repo
	services *services.Service
}

func NewHandler(repo *repo.Repo, service *services.Service) *Handler {
	return &Handler{
		repo:     repo,
		services: service,
	}
}

func (h *Handler) TestAPI(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"time": time.Now().String(),
		"data": h.repo.Stats(),
	})
}
