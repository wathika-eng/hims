package repo

import (
	"context"
	"log"

	"hims/pkg/models"

	"golang.org/x/crypto/bcrypt"
)

func (r *Repo) Seed() error {
	ctx := context.Background()

	count, err := r.db.NewSelect().Model((*models.Doctor)(nil)).Count(ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Seed: data already exists, skipping")
		return nil
	}
	log.Println("Seed: seeding sample data...")

	pass, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	hashedPass := string(pass)

	doctors := []*models.Doctor{
		{FirstName: "Demo", LastName: "Doctor", Email: "john.doe@hospital.com", Password: hashedPass, LicenseNumber: "LIC-001", Specialization: "Cardiology"},
		{FirstName: "Demo", LastName: "Doctor", Email: "jane.smith@hospital.com", Password: hashedPass, LicenseNumber: "LIC-002", Specialization: "Pediatrics"},
		{FirstName: "Demo", LastName: "Doctor", Email: "samuel.mwangi@hospital.com", Password: hashedPass, LicenseNumber: "LIC-003", Specialization: "Internal Medicine"},
	}
	if _, err := r.db.NewInsert().Model(&doctors).Exec(ctx); err != nil {
		return err
	}

	patients := []*models.Patient{
		{FirstName: "Demo", LastName: "Patient", IDNumber: "DEMO001", PhoneNumber: "254700000001", Gender: "female", Age: 30},
		{FirstName: "Demo", LastName: "Patient", IDNumber: "DEMO002", PhoneNumber: "254700000002", Gender: "male", Age: 45},
		{FirstName: "Demo", LastName: "Patient", IDNumber: "DEMO003", PhoneNumber: "254700000003", Gender: "female", Age: 25},
		{FirstName: "Demo", LastName: "Patient", IDNumber: "DEMO004", PhoneNumber: "254700000004", Gender: "male", Age: 55},
		{FirstName: "Demo", LastName: "Patient", IDNumber: "DEMO005", PhoneNumber: "254700000005", Gender: "female", Age: 35},
		{FirstName: "Demo", LastName: "Patient", IDNumber: "DEMO006", PhoneNumber: "254700000006", Gender: "male", Age: 60},
		{FirstName: "Demo", LastName: "Patient", IDNumber: "DEMO007", PhoneNumber: "254700000007", Gender: "female", Age: 28},
		{FirstName: "Demo", LastName: "Patient", IDNumber: "DEMO008", PhoneNumber: "254700000008", Gender: "male", Age: 40},
		{FirstName: "Demo", LastName: "Patient", IDNumber: "DEMO009", PhoneNumber: "254700000009", Gender: "female", Age: 32},
		{FirstName: "Demo", LastName: "Patient", IDNumber: "DEMO010", PhoneNumber: "254700000010", Gender: "male", Age: 50},
	}
	if _, err := r.db.NewInsert().Model(&patients).Exec(ctx); err != nil {
		return err
	}

	programs := []*models.Program{
		{Name: "Hypertension Management", Code: 101},
		{Name: "Diabetes Care", Code: 102},
		{Name: "Maternal Health", Code: 103},
		{Name: "Child Immunization", Code: 104},
		{Name: "HIV/AIDS Support", Code: 105},
		{Name: "TB Treatment", Code: 106},
		{Name: "Malaria Prevention", Code: 107},
		{Name: "Nutrition Counseling", Code: 108},
	}
	if _, err := r.db.NewInsert().Model(&programs).Exec(ctx); err != nil {
		return err
	}

	enrollments := []*models.PatientProgram{
		{PatientID: 1, ProgramID: 1},
		{PatientID: 1, ProgramID: 3},
		{PatientID: 2, ProgramID: 2},
		{PatientID: 2, ProgramID: 1},
		{PatientID: 3, ProgramID: 4},
		{PatientID: 3, ProgramID: 8},
		{PatientID: 4, ProgramID: 1},
		{PatientID: 4, ProgramID: 2},
		{PatientID: 5, ProgramID: 5},
		{PatientID: 5, ProgramID: 8},
		{PatientID: 6, ProgramID: 6},
		{PatientID: 6, ProgramID: 1},
		{PatientID: 7, ProgramID: 3},
		{PatientID: 8, ProgramID: 2},
		{PatientID: 8, ProgramID: 7},
		{PatientID: 9, ProgramID: 5},
		{PatientID: 9, ProgramID: 4},
		{PatientID: 10, ProgramID: 1},
		{PatientID: 10, ProgramID: 6},
	}
	if _, err := r.db.NewInsert().Model(&enrollments).Exec(ctx); err != nil {
		return err
	}

	log.Println("Seed: sample data inserted successfully")
	return nil
}
