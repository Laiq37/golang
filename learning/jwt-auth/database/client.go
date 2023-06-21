package database

import (
	"jwt-authentication/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

// connect to mysql db
func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

// Migrate function will ensuere that users table exist, if not then it will migrate and create
// table from user model
func Migrate() {
	Instance.AutoMigrate(&models.User{})
}
