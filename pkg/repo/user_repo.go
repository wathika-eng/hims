package repo

import (
	"context"
	"errors"
	"fmt"
	"hims/pkg/models"
	"log"

	"github.com/uptrace/bun"
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

	_, err := r.db.NewSelect().Model(&patients).Order("id ASC").
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return patients, nil
}

func (r *Repo) LookupPatient(phoneNum, idNum string) (*models.Patient, error) {
	log.Printf("phone number: %v, ID Number: %v", phoneNum, idNum)
	var patients models.Patient

	err := r.db.NewSelect().Model(&patients).
		WhereGroup(" OR ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.
				WhereOr("phone_number = ?", phoneNum).
				WhereOr("id_number = ?", idNum)
		}).
		Scan(context.Background())

	if err != nil {
		return nil, errors.New("no patient found")
	}
	return &patients, nil
}

func (r *Repo) UpdatePatient(p *models.Patient) (*models.Patient, error) {
	p, err := r.LookupPatient(p.PhoneNumber, p.IDNumber)
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
