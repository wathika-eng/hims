package handlers

import (
	"hims/pkg/models"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func (h *Handler) Profile(ctx *fiber.Ctx) error {
	phoneNum := ctx.Query("phone")
	id := ctx.Query("id")
	log.Printf("id: %v, phone: %v", id, phoneNum)
	if strings.TrimSpace(phoneNum) == "" && strings.TrimSpace(id) == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": "true",
				"data":  "invalid query param",
			},
		)
	}
	patient, err := h.repo.LookupPatient(phoneNum, id)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": "true",
			"data":  err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"error": "false",
		"data":  patient,
	})
}

func (h *Handler) AddPatientProgram(ctx *fiber.Ctx) error {
	if err := h.verifyDoc(ctx); err != nil {
		return err
	}
	type req struct {
		Program models.Program `json:"program" validate:"required"`
		Patient models.Patient `json:"patient" validate:"required"`
	}
	var reqBody req
	if err := ctx.BodyParser(reqBody); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "true",
			"data":  err.Error(),
		})
	}
	
}

// verify token has role as doctor
func (h *Handler) verifyDoc(ctx *fiber.Ctx) error {
	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	// Get the "role" from claims
	role, ok := claims["role"].(string)
	if !ok {
		return fiber.NewError(fiber.StatusForbidden, "Role not found in token")
	}

	if role != "doctor" {
		return fiber.NewError(fiber.StatusForbidden, "Access denied: doctor role required")
	}
	email, _ := claims["email"].(string)
	_, err := h.repo.LookupDoctor(email, "")
	if err != nil {
		return fiber.NewError(fiber.StatusForbidden, "Access denied")
	}
	return ctx.Next()
}
