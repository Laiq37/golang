package database

import (
	"fiber-gorm-crud-api/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func Connect(connectionString string) {
	var err error

	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database")
}

func Migrate() {
	Instance.AutoMigrate(&entities.Dog{})
	log.Println("DataBase Migration Completed ...")
}
