package handlers

import (
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
			"error": err,
		})
	}
	return ctx.JSON(fiber.Map{
		"data":  "doctor created successfully",
		"error": "false",
	})
}

// endpoint to handle patient creation
func (h *Handler) NewPatient(ctx *fiber.Ctx) error {
	var reqBody models.Patient
	if err := ctx.BodyParser(&reqBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := h.services.Validate(reqBody); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	return ctx.JSON(fiber.Map{
		"data":  "patient created successfully",
		"error": "false",
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
			"error": err.Error(),
		})
	}
	if err := h.services.Validate(reqBody); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	token, err := h.services.LoginUser(reqBody.Email, reqBody.Password)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(token)
}
