package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// generate token and sign it with secret key
func (s *Service) generateToken(email, role string) (string, error) {
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "HIMS API",
		"sub": email,
		"aud": role,
		"iat": time.Now(),
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	token, err := jwt.SignedString([]byte(s.cfg.SECRET_KEY))
	if err != nil {
		return "", fmt.Errorf("error generating access token: %v", err)
	}
	return token, nil
}

// takes in password string and returns the hash using bcrypt
func hashPass(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing the password: %v", err)
	}
	return string(hashedPass), nil
}

// compare user input password and the password on db
func checkPassword(docPass, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(docPass), []byte(password))
}
