package entities

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty" gorm:"unique"`
	Password   string `json:"password,omitempty"`
	Address    string `json:"address"`
	PhoneNo    string `json:"phone_no"`
	Country    string `json:"country"`
	City       string `json:"state"`
	PostalCode string `json:"postal_code"`
	Type       string `json:"type,omitempty"`
	Payments   []Payment
	CartItems  []CartItem
	Orders     []Order
}

func (user *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return errors.New("something went wrong!")
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return errors.New("invalid password!")
	}
	return nil
}
