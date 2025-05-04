package handlers

import (
	"hims/pkg/models"

	"github.com/gofiber/fiber/v2"
)

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

func (h *Handler) AddPatientProgram(ctx *fiber.Ctx) error {
	// if err := h.verifyDoc(ctx); err != nil {
	// 	return err
	// }
	type req struct {
		Program string `json:"programName" validate:"required"`
		Patient string `json:"patientID" validate:"required"`
	}
	var reqBody req
	if err := ctx.BodyParser(&reqBody); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}
	if err := h.services.Validate(reqBody); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err,
		})
	}
	program, err := h.repo.LookupProgram(reqBody.Program, 0)
	if err != nil || program.Name == "" {
		return ctx.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}

	patient, err := h.repo.LookupPatient("", reqBody.Patient)
	if err != nil || patient.IDNumber == "" {
		return ctx.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}

	updatedPatient, err := h.services.ModPatient(&patient, &program)
	if err != nil || patient.IDNumber == "" {
		return ctx.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"error": false,
		"data":  updatedPatient,
		//fmt.Sprintf("patient with ID: %v added to program: %v",
		//	updatedPatient.ID, updatedPatient.Programs),
	})
}

func (h *Handler) GetPrograms(ctx *fiber.Ctx) error {
	programs, err := h.repo.FetchPrograms()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}
	if len(programs) == 0 {
		return ctx.Status(204).JSON(fiber.Map{
			"error": false,
			"data":  []models.Patient{},
		})
	}
	return ctx.JSON(fiber.Map{
		"error": false,
		"data":  programs,
	})
}

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
			"data":  []models.Program{},
		})
	}
	return ctx.JSON(fiber.Map{
		"error": false,
		"data":  patients,
	})
}

// // verify token has role as doctor
// func (h *Handler) verifyDoc(ctx *fiber.Ctx) error {
// 	user := ctx.Locals("jwt").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	// Get the "role" from claims
// 	role, ok := claims["aud"].(string)
// 	if !ok {
// 		return fiber.NewError(fiber.StatusForbidden, "Role not found in token")
// 	}
// 	log.Println(role)
// 	if role != "doctor" {
// 		return fiber.NewError(fiber.StatusForbidden, "Access denied: doctor role required")
// 	}
// 	// email, _ := claims["sub"].(string)
// 	// _, err := h.repo.LookupDoctor(email, "")
// 	// if err != nil {
// 	// 	return fiber.NewError(fiber.StatusForbidden, "Access denied")
// 	// }
// 	return ctx.Next()
// }
