package utils

import (
	"errors"
	"regexp"
	"golang.org/x/crypto/bcrypt"
)

// ValidatePassword validates the password according to predefined rules
func ValidatePassword(password string) error {
	// Define regex patterns
	var (
		minLength      = 8
		hasUppercase   = `[A-Z]`
		hasLowercase   = `[a-z]`
		hasDigit       = `[0-9]`
		hasSpecialChar = `[@#$%^&*(),.?":{}|<>]`
	)

	if len(password) < minLength {
		return errors.New("password must be at least 8 characters long")
	}
	if match, _ := regexp.MatchString(hasUppercase, password); !match {
		return errors.New("password must contain at least one uppercase letter")
	}
	if match, _ := regexp.MatchString(hasLowercase, password); !match {
		return errors.New("password must contain at least one lowercase letter")
	}
	if match, _ := regexp.MatchString(hasDigit, password); !match {
		return errors.New("password must contain at least one digit")
	}
	if match, _ := regexp.MatchString(hasSpecialChar, password); !match {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

// HashPassword hashes a plain-text password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}