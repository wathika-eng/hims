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
		{FirstName: "John", LastName: "Doe", Email: "john.doe@hospital.com", Password: hashedPass, LicenseNumber: "LIC-001", Specialization: "Cardiology"},
		{FirstName: "Jane", LastName: "Smith", Email: "jane.smith@hospital.com", Password: hashedPass, LicenseNumber: "LIC-002", Specialization: "Pediatrics"},
		{FirstName: "Samuel", LastName: "Mwangi", Email: "samuel.mwangi@hospital.com", Password: hashedPass, LicenseNumber: "LIC-003", Specialization: "Internal Medicine"},
	}
	if _, err := r.db.NewInsert().Model(&doctors).Exec(ctx); err != nil {
		return err
	}

	patients := []*models.Patient{
		{FirstName: "Alice", LastName: "Wanjiku", IDNumber: "12345678", PhoneNumber: "254712345678", Gender: "female", Age: 30},
		{FirstName: "Bob", LastName: "Kiplagat", IDNumber: "23456789", PhoneNumber: "254723456789", Gender: "male", Age: 45},
		{FirstName: "Carol", LastName: "Akinyi", IDNumber: "34567890", PhoneNumber: "254734567890", Gender: "female", Age: 25},
		{FirstName: "David", LastName: "Omondi", IDNumber: "45678901", PhoneNumber: "254745678901", Gender: "male", Age: 55},
		{FirstName: "Esther", LastName: "Muthoni", IDNumber: "56789012", PhoneNumber: "254756789012", Gender: "female", Age: 35},
		{FirstName: "Francis", LastName: "Kamau", IDNumber: "67890123", PhoneNumber: "254767890123", Gender: "male", Age: 60},
		{FirstName: "Grace", LastName: "Nyambura", IDNumber: "78901234", PhoneNumber: "254778901234", Gender: "female", Age: 28},
		{FirstName: "Henry", LastName: "Kiprop", IDNumber: "89012345", PhoneNumber: "254789012345", Gender: "male", Age: 40},
		{FirstName: "Irene", LastName: "Chebet", IDNumber: "90123456", PhoneNumber: "254790123456", Gender: "female", Age: 32},
		{FirstName: "James", LastName: "Otieno", IDNumber: "01234567", PhoneNumber: "254701234567", Gender: "male", Age: 50},
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
