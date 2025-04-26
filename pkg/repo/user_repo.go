package repo

import (
	"context"
	"fmt"
	"hims/pkg/models"
)

func (r *Repo) InsertNewPatient(p *models.Patient) error {
	_, err := r.db.NewInsert().Model(&p).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) FetchPatients() ([]models.Patient, error) {
	var patients []models.Patient

	_, err := r.db.NewSelect().Model(&patients).Order("id ASC").
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return patients, nil
}

func (r *Repo) LookupPatient(p *models.Patient) (*models.Patient, error) {
	err := r.db.NewSelect().Model(&p).WherePK().Scan(context.Background())
	if err != nil {
		return nil, fmt.Errorf("no patient found with ID: %v", p.ID)
	}
	return p, nil
}

func (r *Repo) UpdatePatient(p *models.Patient) (*models.Patient, error) {
	p, err := r.LookupPatient(p)
	if err != nil {
		return nil, err
	}
	resp, err := r.db.NewUpdate().Model(p).WherePK().Exec(context.Background())
	if err != nil {
		return nil, err
	}
	rowsAffected, _ := resp.RowsAffected()
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no patient found with ID %d", p.ID)
	}

	return p, nil
}
