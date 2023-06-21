package database

import (
	"crud-rest-with-mux/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Panicln("Connected to Database")
}

func Migrate() {
	Instance.AutoMigrate(&models.Product{})
	log.Println("DataBase Migration Completed ...")
}
