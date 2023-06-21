package database

import (
	"ecommerce-store-service/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func Connect(connectString string) {
	var err error

	Instance, err = gorm.Open(mysql.Open(connectString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database")
}

func Migrate() {
	log.Println("Migration Database ....")
	Instance.AutoMigrate(&entities.ProductCategory{}, &entities.Product{}, &entities.User{}, &entities.Payment{}, &entities.Order{}, &entities.OrderItem{}, &entities.CartItem{})
	log.Println("Database Mirgation has been Completed!")
}
