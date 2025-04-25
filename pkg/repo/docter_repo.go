package repo

import (
	"context"
	"fmt"
	"hims/pkg/models"
)

func (r *Repo) InsertNewDoctor(d *models.Doctor) error {
	_, err := r.db.NewInsert().Model(d).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) InsertNewProgram(p *models.Program) error {
	_, err := r.db.NewInsert().Model(p).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) FetchPrograms() ([]models.Program, error) {
	var programs []models.Program

	_, err := r.db.NewSelect().Model(&programs).Order("id ASC").
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return programs, nil
}

func (r *Repo) LookupProgram(p *models.Program) (*models.Program, error) {
	err := r.db.NewSelect().Model(p).Where("name = ?", p.Name).Scan(context.Background())
	if err != nil {
		return nil, fmt.Errorf("no program found with name: %v", p.ID)
	}
	return p, nil
}
