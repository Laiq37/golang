package handlers

import (
	"ecommerce-store-service/auth"
	"ecommerce-store-service/database"
	"ecommerce-store-service/entities"
	"ecommerce-store-service/helpers"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) error {
	var user entities.User

	//if body missing
	if len(c.Body()) == 0 {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"message": "body is missing",
			"data":    nil,
		})
	}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"message": "invalid data",
			"data":    nil,
		})
	}

	err = helpers.ValidateSignupRequiredFields(user)
	//if required fields are missing
	if err != nil {
		return c.Status(422).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}
	err = helpers.EmailValidation(user.Email)
	//if email not in correct format
	if err != nil {
		return c.Status(422).JSON(&fiber.Map{
			"success": false,
			"message": "invalid email format",
			"data":    nil,
		})
	}
	err = user.HashPassword()
	//if password not hash
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}
	result := database.Instance.Create(&user)
	if result.RowsAffected == 0 {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "user not created",
			"data":    nil,
		})
	}
	return c.Status(201).JSON(&fiber.Map{
		"success": true,
		"message": "user has been created",
		"data":    user,
	})
}

func Login(c *fiber.Ctx) error {
	//if body missing
	if len(c.Body()) == 0 {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"message": "body is missing",
			"data":    nil,
		})
	}
	var request helpers.TokenRequest
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"message": "invalid data",
			"data":    nil,
		})
	}
	err = helpers.ValidateLoginRequiredFields(request)
	//if required fields are missing
	if err != nil {
		return c.Status(422).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}
	var dbUser *entities.User
	result := database.Instance.Where("email = ?", request.Email).First(&dbUser)
	if result.RowsAffected == 0 {
		return c.Status(422).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}
	err = dbUser.CheckPassword(request.Password)
	if err != nil {
		return c.Status(422).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}
	tokenString, expiresAt, err := auth.GenerateJWT(dbUser.Name, dbUser.Email, dbUser.Type)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Login Successful",
		"data": &fiber.Map{
			"token":     tokenString,
			"expiresAt": expiresAt,
		},
	})
}
