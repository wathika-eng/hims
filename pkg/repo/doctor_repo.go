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

func (r *Repo) LookupDoctor(email, specialization string) (models.Doctor, error) {
	var doctor models.Doctor
	err := r.db.NewSelect().
		Model(&doctor).WhereOr("email = ?", email).
		WhereOr("specialization = ?", specialization).Limit(1).
		Scan(context.Background())
	if err != nil {
		return models.Doctor{}, fmt.Errorf("no doctor found with email: %v", email)
	}
	return doctor, nil
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

	err := r.db.NewSelect().Model(&programs).Order("id ASC").
		Scan(context.Background())
	if err != nil {
		return []models.Program{}, err
	}
	return programs, nil
}

func (r *Repo) LookupProgram(name string, code int) (models.Program, error) {
	var program models.Program

	err := r.db.NewSelect().Model(program).
		WhereOr("name = ?", name).WhereOr("code = ?", code).Limit(1).
		Scan(context.Background())
	if err != nil {
		return models.Program{}, fmt.Errorf("no program found with name: %v", name)
	}
	return program, nil
}
