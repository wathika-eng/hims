package services

import (
	"fmt"
	"hims/pkg/models"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateDoctor(d *models.Doctor) error {
	hashedPass, err := hashPass(d.Password)
	if err != nil {
		return err
	}
	d.Password = hashedPass

	return s.repo.InsertNewDoctor(d)
}

func (s *Service) LoginUser(email, password string) (string, error) {
	doctor, err := s.repo.LookupDoctor(email, "")
	if err != nil {
		return "", err
	}
	if err := checkPassword(doctor.Password, password); err != nil {
		return "", fmt.Errorf("wrong email and/or password")
	}
	return "", nil
}

func hashPass(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing the password: %v", err)
	}
	return string(hashedPass), nil
}

func checkPassword(docPass, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(docPass), []byte(password))
}
