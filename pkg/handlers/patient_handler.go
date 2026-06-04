package handlers

import (
	"fmt"
	"hims/pkg/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// endpoint to handle patient creation
func (h *Handler) NewPatient(ctx fiber.Ctx) error {
	var reqBody models.Patient
	if err := ctx.Bind().Body(&reqBody); err != nil {
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

// /api/v1/protected/patients/:id
func (h *Handler) Profile(ctx fiber.Ctx) error {
	rawID := ctx.Params("id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": true,
			"data":  "invalid patient id",
		})
	}

	patient, err := h.repo.GetPatientByID(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
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
func (h *Handler) GetPatients(ctx fiber.Ctx) error {
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
