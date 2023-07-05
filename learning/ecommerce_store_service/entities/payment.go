package entities

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	OrderID uint
	Time    int64
}
