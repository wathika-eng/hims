package services

import (
	"fmt"
	"hims/pkg/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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
	return "", nil
}

func generateToken() (string, error){
	jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":
		"sub":
		"iat":
	})
}

// takes in password string and returns the hash using bcrypt
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
