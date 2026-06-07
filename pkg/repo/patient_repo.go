package repo

import (
	"context"
	"errors"
	"hims/pkg/models"
	"log"
	"time"

	"github.com/uptrace/bun/driver/pgdriver"
)

func (r *Repo) InsertNewPatient(p *models.Patient) error {
	_, err := r.db.NewInsert().Model(p).Exec(context.Background())
	if err != nil {
		if err, ok := err.(pgdriver.Error); ok && err.IntegrityViolation() {
			return errors.New("patient with this id number or email already exists")
		}
		return err
	}
	return nil
}

func (r *Repo) FetchPatients() ([]models.Patient, error) {
	var patients []models.Patient

	err := r.db.NewSelect().Model(&patients).Relation("Programs").
		Scan(context.Background())
	if err != nil {
		return []models.Patient{}, err
	}

	return patients, nil
}

func (r *Repo) LookupPatient(phoneNum, idNum string) (models.Patient, error) {
	var patient models.Patient

	q := r.db.NewSelect().Model(&patient).Relation("Programs")
	if idNum != "" {
		q = q.Where("id_number = ?", idNum)
	}
	if phoneNum != "" {
		q = q.Where("phone_number = ?", phoneNum)
	}
	err := q.Limit(1).Scan(context.Background())
	if err != nil {
		return models.Patient{}, errors.New("patient not found")
	}

	return patient, nil
}

func (r *Repo) GetPatientByID(id int) (models.Patient, error) {
	var patient models.Patient
	err := r.db.NewSelect().Model(&patient).Relation("Programs").
		Where("id = ?", id).Limit(1).Scan(context.Background())
	if err != nil {
		return models.Patient{}, errors.New("patient not found")
	}
	return patient, nil
}

func (r *Repo) UpdatePatient(p *models.Patient, program *models.Program) (*models.Patient, error) {
	_, err := r.LookupPatient(p.PhoneNumber, p.IDNumber)

	if err != nil {
		return nil, err
	}

	p.UpdatedAt.Time = time.Now()
	_, err = r.db.NewUpdate().Model(p).
		WherePK().Exec(context.Background())

	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error while updating patient")
	}

	patientProgram := models.PatientProgram{
		PatientID: p.ID,
		Patient:   p,
		ProgramID: program.ID,
		Program:   program,
	}
	_, err = r.db.NewInsert().Model(&patientProgram).Exec(context.Background())
	if err != nil {
		return nil, errors.New("error while inserting new program data to patient, duplicate")
	}
	return p, nil
}
