package handlers

import (
	"fmt"
	"hims/pkg/models"

	"github.com/gofiber/fiber/v2"
)

// endpoint to handle doc creation
func (h *Handler) NewDoctor(ctx *fiber.Ctx) error {
	var reqBody models.Doctor

	if err := ctx.BodyParser(&reqBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.services.Validate(reqBody); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": true,
			"data":  err,
		})
	}

	if err := h.services.CreateDoctor(&reqBody); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"data":  fmt.Sprintf("doctor %v created successfully", reqBody.Email),
		"error": false,
	})
}

func (h *Handler) LoginDoctor(ctx *fiber.Ctx) error {
	type req struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	var reqBody req

	if err := ctx.BodyParser(&reqBody); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}
	if err := h.services.Validate(reqBody); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": true,
			"data":  err,
		})
	}
	token, err := h.services.LoginUser(reqBody.Email, reqBody.Password)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"accessToken": token,
	})
}
