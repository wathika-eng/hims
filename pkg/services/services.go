// package services handles logic and validation checks
package services

import (
	"hims/pkg/repo"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Service struct {
	repo *repo.Repo
}

func NewServices(repo *repo.Repo) *Service {
	return &Service{
		repo: repo,
	}
}
