package handlers

import (
	"fiber-gorm-crud-api/database"
	"fiber-gorm-crud-api/entities"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetDogs(c *fiber.Ctx) error {
	var dogs []entities.Dog

	database.Instance.Find(&dogs)
	log.Println(dogs)
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"data":    &dogs,
	})
}

func GetDog(c *fiber.Ctx) error {
	id := c.Params("id")
	var dog entities.Dog

	result := database.Instance.Find(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(dog)
}

func AddDog(c *fiber.Ctx) error {
	var dog entities.Dog

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.Instance.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	dog := new(entities.Dog)
	id := c.Params("id")

	if err := c.BodyParser(dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.Instance.Where("id= ?", id).Updates(&dog)
	return c.Status(200).JSON(&dog)
}

func RemoveDog(c *fiber.Ctx) error {
	id := c.Params("id")
	var dog entities.Dog
	result := database.Instance.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}
