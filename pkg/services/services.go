// package services handles logic and validation checks
package services

import "hims/pkg/repo"

type Service struct {
	repo *repo.Repo
}

func NewServices(repo *repo.Repo) *Service {
	return &Service{
		repo: repo,
	}
}
