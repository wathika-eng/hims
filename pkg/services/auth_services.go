package services

import (
	"fmt"
	"hims/pkg/models"
	"log"
)

// calls the insert function on repo, if validate checks pass
func (s *Service) CreateDoctor(d *models.Doctor) error {
	d.Role = ""
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

func (s *Service) CreatePatient(p *models.Patient) error {
	cleanNum, err := cleanUpPhone(p.PhoneNumber)
	if err != nil {
		return err
	}
	p.PhoneNumber = cleanNum
	p.Role = ""
	return s.repo.InsertNewPatient(p)
}

func (s *Service) ModPatient(p *models.Patient, program *models.Program) (*models.Patient, error) {
	p.Programs = append(p.Programs, program)
	patient, err := s.repo.UpdatePatient(p, program)
	log.Println(err.Error())
	if err != nil {
		return nil, err
	}

	return patient, nil
}
