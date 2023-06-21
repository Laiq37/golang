package entities

import (
	"time"
)

type Order struct {
	ID         uint `gorm:"primary_key"`
	UserID     uint
	OrderDate  time.Time
	OrderItems []OrderItem
	Payments   []Payment
	Status     string
}
