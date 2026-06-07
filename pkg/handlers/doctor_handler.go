package handlers

import (
	"fmt"
	"hims/pkg/models"

	"github.com/gofiber/fiber/v3"
)

// endpoint to handle doc creation
func (h *Handler) NewDoctor(ctx fiber.Ctx) error {
	var input models.DoctorSignupInput

	if err := ctx.Bind().Body(&input); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": true,
			"data":  "invalid request body",
		})
	}

	if err := h.services.Validate(input); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": true,
			"data":  err,
		})
	}

	doctor := models.Doctor{
		FirstName:      input.FirstName,
		LastName:       input.LastName,
		Email:          input.Email,
		Password:       input.Password,
		LicenseNumber:  input.LicenseNumber,
		Specialization: input.Specialization,
	}

	if err := h.services.CreateDoctor(&doctor); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  "failed to create account",
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"data":  fmt.Sprintf("doctor %v created successfully", input.Email),
		"error": false,
	})
}

func (h *Handler) LoginDoctor(ctx fiber.Ctx) error {
	type req struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	var reqBody req

	if err := ctx.Bind().Body(&reqBody); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": true,
			"data":  "invalid request body",
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
		return ctx.Status(401).JSON(fiber.Map{
			"error": true,
			"data":  "invalid email or password",
		})
	}
	return ctx.JSON(fiber.Map{
		"accessToken": token,
	})
}
