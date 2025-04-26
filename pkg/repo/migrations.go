package repo

import (
	"context"
	"fmt"
	"hims/pkg/models"
	"log"
)

func (r *Repo) Up() error {
	ctx := context.Background()
	r.db.RegisterModel((*models.PatientProgram)(nil))
	modelsToMigrate := []any{
		(*models.Doctor)(nil),
		(*models.PatientProgram)(nil),
		(*models.Patient)(nil),
		(*models.Program)(nil),
	}

	for _, model := range modelsToMigrate {
		_, err := r.db.NewCreateTable().
			Model(model).
			IfNotExists().
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("migration failed for model %T: %v", model, err)
		}
	}
	log.Println("Tables created (if not exists)")
	return nil
}

func (r *Repo) Reset() error {
	ctx := context.Background()

	modelsToReset := []any{
		(*models.Doctor)(nil),
		(*models.Patient)(nil),
		(*models.Program)(nil),
	}

	err := r.db.ResetModel(ctx, modelsToReset...)
	if err != nil {
		return fmt.Errorf("reset failed: %v", err)
	}
	log.Println("Tables successfully reset")
	return nil
}
