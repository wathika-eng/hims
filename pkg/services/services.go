// package services handles logic and validation checks
package services

import (
	"hims/pkg/config"
	"hims/pkg/repo"
)

type Service struct {
	repo *repo.Repo
	cfg  *config.Config
}

func NewServices(repo *repo.Repo, cfg *config.Config) *Service {
	return &Service{
		repo: repo,
		cfg:  cfg,
	}
}
