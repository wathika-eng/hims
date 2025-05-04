package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// generate token and sign it with secret key
func (s *Service) generateToken(email, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "HIMS API",
		"sub": email,
		"aud": role,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	signedToken, err := token.SignedString([]byte(s.cfg.SECRET_KEY))
	if err != nil {
		return "", fmt.Errorf("error generating access token: %v", err)
	}
	return signedToken, nil
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

func cleanUpPhone(num string) (string, error) {
	var cleanedNum string
	// remove any spaces
	phoneNum := strings.TrimSpace(num)
	phoneNum = strings.ReplaceAll(phoneNum, "-", "")

	if strings.HasPrefix(phoneNum, "07") {
		cleanedNum = fmt.Sprintf("%v", "254"+phoneNum[1:])
		return cleanedNum, nil
	} else if strings.HasPrefix(phoneNum, "+254") {
		cleanedNum = fmt.Sprintf("%v", "254"+phoneNum[4:])
		return cleanedNum, nil
	} else {
		return "", fmt.Errorf("only kenyan phone numbers accepted :%v", phoneNum)
	}
}
