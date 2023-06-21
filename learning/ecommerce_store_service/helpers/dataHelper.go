package helpers

import (
	"ecommerce-store-service/database"
	"ecommerce-store-service/entities"
)

func GetUserId(email string) (id uint) {
	var user entities.User
	result := database.Instance.Where("email = ?", email).First(&user)
	id = user.ID
	if result.RowsAffected == 0 {
		return id
	}
	return id
}
