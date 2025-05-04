package handlers

import (
	"bytes"
	"fmt"
	"hims/pkg/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/phpdave11/gofpdf"
)

func (h *Handler) AddProgram(ctx *fiber.Ctx) error {
	var reqBody models.Program
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
	if err := h.repo.InsertNewProgram(&reqBody); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}
	return ctx.Status(201).JSON(fiber.Map{
		"error": false,
		"data":  fmt.Sprintf("added %v program", reqBody.Name),
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
			"data":  []models.Program{},
		})
	}

	return ctx.JSON(fiber.Map{
		"error": false,
		"data":  programs,
	})
}

// returns a pdf of programs and patients enrolled in each
func (h *Handler) GeneratePDF(ctx *fiber.Ctx) error {
	programs, err := h.repo.FetchPrograms()
	if err != nil {
		return ctx.Status(500).SendString("Failed to fetch programs")
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "", 12)

	// watermark
	pdf.SetHeaderFunc(func() {
		pdf.SetFont("Arial", "I", 40)
		pdf.SetTextColor(200, 200, 200)
		for y := 50.0; y < 250; y += 50 {
			pdf.Text(50, y, "CONFIDENTIAL")
		}
		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("Arial", "", 12)
	})

	pdf.AddPage()

	pdf.SetFont("Arial", "B", 20)
	pdf.CellFormat(0, 15, "Programs and Patients Summary", "", 1, "C", false, 0, "")
	pdf.Ln(5)

	for _, prog := range programs {
		pdf.SetFont("Arial", "B", 14)
		pdf.CellFormat(0, 10, "Program: "+prog.Name, "", 1, "", false, 0, "")
		pdf.SetFont("Arial", "", 12)
		pdf.CellFormat(0, 8, fmt.Sprintf("Program Code: %d", prog.Code), "", 1, "", false, 0, "")
		pdf.Ln(2)

		pdf.SetFont("Arial", "B", 12)
		pdf.CellFormat(30, 7, "ID", "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 7, "Name", "1", 0, "", false, 0, "")
		pdf.CellFormat(30, 7, "Phone", "1", 0, "", false, 0, "")
		pdf.CellFormat(20, 7, "Gender", "1", 0, "", false, 0, "")
		pdf.CellFormat(20, 7, "Age", "1", 1, "", false, 0, "")

		pdf.SetFont("Arial", "", 11)
		for _, patient := range prog.Patients {
			fullName := patient.FirstName + " " + patient.LastName
			pdf.CellFormat(30, 7, patient.IDNumber, "1", 0, "", false, 0, "")
			pdf.CellFormat(40, 7, fullName, "1", 0, "", false, 0, "")
			pdf.CellFormat(30, 7, patient.PhoneNumber, "1", 0, "", false, 0, "")
			pdf.CellFormat(20, 7, patient.Gender, "1", 0, "", false, 0, "")
			pdf.CellFormat(20, 7, strconv.Itoa(int(patient.Age)), "1", 1, "", false, 0, "")
		}
		pdf.Ln(5)
	}
	// pages := pdf.PageCount()
	// pdf.SetFooterFunc(func() {
	// 	pdf.SetFont("Arial", "B", 14)
	// 	pdf.Text(50, 50, fmt.Sprintf("page number: %v", strconv.Itoa(pages)))
	// })

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return ctx.Status(500).SendString("Error generating PDF")
	}

	ctx.Set("Content-Type", "application/pdf")
	ctx.Set("Content-Disposition", "attachment; filename=programs_summary.pdf")
	return ctx.SendStream(&buf)
}
