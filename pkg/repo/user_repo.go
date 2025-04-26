package repo

import (
	"context"
	"errors"
	"fmt"
	"hims/pkg/models"
	"log"
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

	query := `SELECT * FROM patients ORDER BY id ASC`
	err := r.db.NewRaw(query).Scan(context.Background(), &patients)
	if err != nil {
		log.Println("Error executing query:", err)
		return []models.Patient{}, err
	}

	return patients, nil
}

func (r *Repo) LookupPatient(phoneNum, idNum string) (models.Patient, error) {
	log.Printf("phone: %v, id: %v", phoneNum, idNum)
	var patient models.Patient

	q := r.db.NewSelect().Model(&patient)
	if phoneNum != "" {
		q = q.Where("phone_number = ?", phoneNum)
	} else if idNum != "" {
		q = q.Where("id_number = ?", idNum)
	} else {
		return models.Patient{}, errors.New("no search parameters provided")
	}

	err := q.Limit(1).Scan(context.Background())
	if err != nil {
		log.Println(err.Error())
		return models.Patient{}, fmt.Errorf("failed to find patient: %w", err)
	}

	return patient, nil
}

func (r *Repo) UpdatePatient(p *models.Patient) (*models.Patient, error) {
	_, err := r.LookupPatient(p.PhoneNumber, p.IDNumber)
	if err != nil {
		return nil, err
	}
	resp, err := r.db.NewUpdate().Model(p).Set("programs = ?", p.Programs).
		WherePK().Exec(context.Background())
	if err != nil {
		return nil, err
	}
	rowsAffected, _ := resp.RowsAffected()
	if rowsAffected == 0 {
		return nil, errors.New("update not done")
	}

	return p, nil
}
