package handlers

import (
	"ecommerce-store-service/auth"
	"ecommerce-store-service/database"
	"ecommerce-store-service/entities"
	"ecommerce-store-service/helpers"

	"github.com/gofiber/fiber/v2"
)

type cartProduct struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Quantity uint   `json:"quantity"`
}

func GetCartItems(c *fiber.Ctx) error {
	email := c.Locals("user").(auth.ValidUser).Email
	id := helpers.GetUserId(email)
	if id == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "failed to get cart items",
			"data":    nil,
		})
	}
	var cartProduct []cartProduct
	result := database.Instance.Raw("SELECT products.id,products.name,cart_items.quantity from products JOIN cart_items ON cart_items.product_id = products.id WHERE cart_items.user_id = ?", id).Scan(&cartProduct)
	if result.RowsAffected == 0 {
		if result.Error != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": true,
				"message": "failed to get cart items",
				"data":    nil,
			})
		}
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "cart is empty!",
			"data":    nil,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Cart item has been successfully fetched",
		"data":    cartProduct,
	})
}

func AddItemToCart(c *fiber.Ctx) error {
	if len(c.Body()) == 0 {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "body is missing",
			"data":    nil,
		})
	}
	email := c.Locals("user").(auth.ValidUser).Email
	id := helpers.GetUserId(email)
	if id == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "item not added to cart",
			"data":    nil,
		})
	}
	var cartITem entities.CartItem
	err := c.BodyParser(&cartITem)
	if err != nil {
		return c.Status(422).JSON(&fiber.Map{
			"success": false,
			"message": "failed to parse due to incorrect data",
			"data":    nil,
		})
	}
	cartITem.UserID = id
	result := database.Instance.Create(&cartITem)
	if result.RowsAffected == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "failed to add product in cart",
			"data":    nil,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Item added to cart!",
		"data":    cartITem,
	})
}

func DeleteCartItem(c *fiber.Ctx) error {
	productId := c.Params("id")
	email := c.Locals("user").(auth.ValidUser).Email
	userId := helpers.GetUserId(email)
	if userId == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "failed to remove item from cart",
			"data":    nil,
		})
	}
	var cartItem entities.CartItem
	result := database.Instance.Where("user_id = ? AND product_id = ?", userId, productId).Delete(&cartItem)
	if result.RowsAffected == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "failed to remove item from cart",
			"data":    nil,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Item removed from cart",
		"data":    nil,
	})
}

func UpdateCart(c *fiber.Ctx) error {
	if len(c.Body()) == 0 {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "body is missing",
			"data":    nil,
		})
	}
	email := c.Locals("user").(auth.ValidUser).Email
	userId := helpers.GetUserId(email)
	if userId == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "failed to update cart",
			"data":    nil,
		})
	}
	var cartItems []entities.CartItem
	if err := c.BodyParser(&cartItems); err != nil {
		return c.Status(422).JSON(&fiber.Map{
			"success": false,
			"message": "field types are incorrects",
			"data":    nil,
		})
	}
	for i := range cartItems {
		cartItems[i].UserID = userId
	}
	result := database.Instance.Where("user_id = ?", userId).Delete(entities.CartItem{})
	if result.RowsAffected == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "failed to update cart!",
			"data":    nil,
		})
	}
	result = database.Instance.Create(&cartItems)
	if result.RowsAffected == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "failed to update cart!",
			"data":    nil,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Cart has been updated successfully",
		"data":    cartItems,
	})
}
