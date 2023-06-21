package helpers

import (
	"ecommerce-store-service/entities"
	"errors"
	"net/mail"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func EmailValidation(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

func ValidateSignupRequiredFields(user entities.User) error {
	if user.Email == "" {
		return errors.New("Email is Required")
	}
	if user.Password == "" {
		return errors.New("Password is Required")
	}
	if user.Type == "" {
		return errors.New("Type is Required")
	}
	return nil
}

func ValidateLoginRequiredFields(user TokenRequest) error {
	if user.Email == "" {
		return errors.New("Email is Required")
	}
	if user.Password == "" {
		return errors.New("Password is Required")
	}
	return nil
}
