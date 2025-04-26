package services

import (
	"fmt"
	"hims/pkg/models"
)

// calls the insert function on repo, if validate checks pass
func (s *Service) CreateDoctor(d *models.Doctor) error {
	hashedPass, err := hashPass(d.Password)
	if err != nil {
		return err
	}
	d.Password = hashedPass

	return s.repo.InsertNewDoctor(d)
}

// returns a jwt token if user is in the db
func (s *Service) LoginUser(email, password string) (string, error) {
	doctor, err := s.repo.LookupDoctor(email, "")
	if err != nil {
		return "", err
	}

	if err := checkPassword(doctor.Password, password); err != nil {
		return "", fmt.Errorf("wrong email and/or password")
	}
	token, err := s.generateToken(doctor.Email, doctor.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}
