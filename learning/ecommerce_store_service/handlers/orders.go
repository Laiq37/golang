package handlers

import (
	"ecommerce-store-service/auth"
	"ecommerce-store-service/database"
	"ecommerce-store-service/entities"
	"ecommerce-store-service/helpers"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type OrderResponse struct {
	ID         uint                 `json:"id"`
	OrderTime  int64                `json:"orderTime"`
	OrderItems []entities.OrderItem `json:"orderItems" gorm:"foreignKey:OrderID"`
	Status     string               `json:"status"`
}

func CreateOrder(c *fiber.Ctx) error {
	email := c.Locals("user").(auth.ValidUser).Email
	userId := helpers.GetUserId(email)
	if userId == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "failed to order",
			"data":    nil,
		})
	}
	if len(c.Body()) == 0 {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "body is missing",
			"data":    nil,
		})
	}
	var order entities.Order
	err := c.BodyParser(&order)
	if err != nil {
		return c.Status(422).JSON(&fiber.Map{
			"success": false,
			"message": "invalid data format",
			"data":    nil,
		})
	}
	var prodIds []string
	for _, item := range order.OrderItems {
		prodIds = append(prodIds, fmt.Sprint(item.ProductID))
	}
	ids := strings.Join(prodIds, ",")
	var products []entities.Product
	result := database.Instance.Where(fmt.Sprintf("id in (%s)", ids)).Find(&products)
	if result.RowsAffected == 0 || len(order.OrderItems) != len(products) {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "failed to get Order items",
			"data":    nil,
		})
	}
	productMap := make(map[uint]string)
	for _, product := range products {
		productMap[product.ID] = product.Name
	}
	for i := range order.OrderItems {
		order.OrderItems[i].Name = productMap[order.OrderItems[i].ProductID]
	}
	order.Payment = entities.Payment{Time: time.Now().Local().Unix()}
	order.UserID = userId
	order.OrderTime = time.Now().Local().Unix()
	order.Status = "pending"
	result = database.Instance.Create(&order)
	if result.RowsAffected == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "failed to generate Order",
			"data":    nil,
		})
	}
	return c.Status(201).JSON(&fiber.Map{
		"success": true,
		"message": "Order has been placed successfully",
		"data": &fiber.Map{
			"id":         order.ID,
			"totalItems": len(order.OrderItems),
			"orderTime":  order.OrderTime,
		},
	})
}

func GetOrders(c *fiber.Ctx) error {
	email := c.Locals("user").(auth.ValidUser).Email
	userId := helpers.GetUserId(email)
	if userId == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "failed to order",
			"data":    nil,
		})
	}
	// var orderResponses []OrderResponse
	var orders []entities.Order
	result :=
		// database.Instance.Table("orders").
		// Select("id, order_time, status").
		// Preload("OrderItems").
		// database.Instance.Table("orders").
		// 	Joins("JOIN order_items ON orders.id = order_items.order_id").
		// 	Where("orders.user_id = ?", 6).Find(&orderResponses)
		database.Instance.Preload("OrderItems").Where("user_id = ?", userId).Find(&orders)
		// SELECT id, status, order_items.quantity, order_items.name, order_items.product_id FROM `orders` JOIN order_items ON orders.id = order_items.order_id WHERE orders.user_id = 6 GROUP BY id
	if result.RowsAffected == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "failed to retrieve orders",
			"data":    nil,
		})
	}
	return c.Status(201).JSON(&fiber.Map{
		"success": true,
		"message": "Order has been placed successfully",
		"data":    orders,
	})
}
