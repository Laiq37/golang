package entities

import "gorm.io/gorm"

type Dog struct {
	gorm.Model
	Name      string `json:"name"`
	Breed     string `json:"breed"`
	Age       string `json:"age"`
	IsGoodBoy bool   `json:"isGoodBoy" gorm:"default:true"`
}
