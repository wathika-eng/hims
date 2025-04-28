package repo

import (
	"context"
	"errors"
	"fmt"
	"hims/pkg/models"
	"log"
	"time"
)

func (r *Repo) InsertNewPatient(p *models.Patient) error {
	_, err := r.db.NewInsert().Model(p).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) FetchPatients() ([]models.Patient, error) {
	var patients []models.Patient

	err := r.db.NewSelect().Model(&patients).Relation("Programs").
		Scan(context.Background())
	if err != nil {
		log.Println("Error executing query:", err)
		return []models.Patient{}, err
	}

	return patients, nil
}

func (r *Repo) LookupPatient(phoneNum, idNum string) (models.Patient, error) {
	log.Printf("phone: %v, id: %v", phoneNum, idNum)
	var patient models.Patient

	q := r.db.NewSelect().Model(&patient).Relation("Programs").
		WhereOr("id_number = ?", idNum).WhereOr("phone_number = ?", phoneNum)
	err := q.Limit(1).Scan(context.Background())
	if err != nil {
		log.Println(err.Error())
		return models.Patient{}, fmt.Errorf("failed to find patient: %w", err)
	}

	return patient, nil
}

func (r *Repo) UpdatePatient(p *models.Patient, program *models.Program) (*models.Patient, error) {
	_, err := r.LookupPatient(p.PhoneNumber, p.IDNumber)
	if err != nil {
		return nil, err
	}

	p.UpdatedAt.Time = time.Now()
	resp, err := r.db.NewUpdate().Model(p).
		WherePK().Exec(context.Background())
	if err != nil {
		return nil, err
	}
	rowsAffected, _ := resp.RowsAffected()
	if rowsAffected == 0 {
		return nil, errors.New("update not done")
	}
	patientProgram := models.PatientProgram{
		PatientID: p.ID,
		Patient:   p,
		ProgramID: program.ID,
		Program:   program,
	}
	_, err = r.db.NewInsert().Model(&patientProgram).Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return p, nil
}
