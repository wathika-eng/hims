package handlers

import (
	"fmt"
	"hims/pkg/models"

	"github.com/gofiber/fiber/v2"
)

// endpoint to handle patient creation
func (h *Handler) NewPatient(ctx *fiber.Ctx) error {
	var reqBody models.Patient
	if err := ctx.BodyParser(&reqBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := h.services.Validate(reqBody); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": true,
			"date":  err,
		})
	}
	if err := h.services.CreatePatient(&reqBody); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}
	return ctx.Status(201).JSON(fiber.Map{
		"data":  fmt.Sprintf("patient %v %v successfully created", reqBody.FirstName, reqBody.LastName),
		"error": false,
	})
}

// api/v1/protected/patients/123
func (h *Handler) Profile(ctx *fiber.Ctx) error {
	param := struct {
		ID    string `params:"id" validate:"max=10"`
		Phone string `params:"phone" validate:"max=10"`
	}{}

	if err := ctx.ParamsParser(&param); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": true,
			"data":  "invalid param, either phone number or id is needed",
		})
	}

	if err := h.services.Validate(&param); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": true,
			"data":  err,
		})
	}

	patient, err := h.repo.LookupPatient(param.Phone, param.ID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"error": false,
		"data":  patient,
	})
}

// get all patients from the db
func (h *Handler) GetPatients(ctx *fiber.Ctx) error {
	patients, err := h.repo.FetchPatients()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}
	if len(patients) == 0 {
		return ctx.Status(204).JSON(fiber.Map{
			"error": false,
			"data":  []models.Patient{},
		})
	}
	return ctx.JSON(fiber.Map{
		"error": false,
		"data":  patients,
	})
}
