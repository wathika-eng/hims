// package services handles logic and validation checks
package services

import (
	"hims/pkg/config"
	"hims/pkg/repo"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	repo      *repo.Repo
	cfg       *config.Config
	Validator *validator.Validate
}

func NewServices(repo *repo.Repo, cfg *config.Config) *Service {
	return &Service{
		repo:      repo,
		cfg:       cfg,
		Validator: validator.New(),
	}
}
