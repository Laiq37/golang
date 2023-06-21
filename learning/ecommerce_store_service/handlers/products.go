package handlers

import (
	"ecommerce-store-service/auth"
	"ecommerce-store-service/database"
	"ecommerce-store-service/entities"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateProduct(c *fiber.Ctx) error {
	if c.Locals("user").(auth.ValidUser).Email != "admin" {
		return c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Not authorized for this action",
			"data":    nil,
		})
	}
	if c.Body() == nil {
		return c.Status(422).JSON(&fiber.Map{
			"success": false,
			"message": "body is missing",
			"data":    nil,
		})
	}
	var product entities.ProductCategory
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(422).JSON(&fiber.Map{
			"success": false,
			"message": "fields are missing",
			"data":    nil,
		})
	}
	result := database.Instance.Where("name= ?", product.Name).Find(&product)
	if result.RowsAffected == 0 {
		result = database.Instance.Create(&product)
	} else {
		for i, _ := range product.Products {
			product.Products[i].ProductCategoryID = product.ID
		}
		result = database.Instance.Create(&product.Products)
	}
	if result.RowsAffected == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": result.Error,
			"data":    nil,
		})
	}
	return c.Status(201).JSON(&fiber.Map{
		"success": true,
		"message": "product has been successfully added",
		"data":    product,
	})
}

func GetProducts(c *fiber.Ctx) error {
	categoryNames := c.Query("category")
	var products []entities.Product
	var productCategories []entities.ProductCategory
	var result *gorm.DB
	if categoryNames == "" {
		result = database.Instance.Find(&products)
	} else {
		regex := regexp.MustCompile(`\[(.*?)\]`)
		categoryNames = regex.ReplaceAllString(categoryNames, "($1)")
		query := fmt.Sprintf("name IN %s", categoryNames)
		result = database.Instance.Preload("Products").Where(query).Find(&productCategories)
		// result = database.Instance.Joins("JOIN product_categories ON product_categories.id = products.product_category_id").Where("product_categories.name = ?", categoryNames).Find(&products)
		// .Preload("ProductCategory").Where("name = ?", categoryNames).Find()
	}
	if result.RowsAffected == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "failed to get product",
			"data":    nil,
		})
	}
	if len(products) > 0 {
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "All product has been fetched successfully",
			"data":    &products,
		})
	}
	for i := range productCategories {
		products = append(products, productCategories[i].Products...)
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "All product has been fetched successfully",
		"data":    products,
	})
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product entities.Product
	result := database.Instance.Find(&product, id)
	if result.RowsAffected == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "failed to get product",
			"data":    nil,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Product has been fetched successfully",
		"data":    product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	if c.Locals("user").(auth.ValidUser).Email != "admin" {
		return c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Not authorized for this action",
			"data":    nil,
		})
	}
	id := c.Params("id")
	var product entities.Product
	result := database.Instance.Delete(&product, id)
	if result.RowsAffected == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "failed to delete product",
			"data":    nil,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Product has been deleted successfully",
		"data":    product,
	})
}

func DeleteProducts(c *fiber.Ctx) error {
	if c.Locals("user").(auth.ValidUser).Email != "admin" {
		return c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Not authorized for this action",
			"data":    nil,
		})
	}
	ids := c.Query("prodIds")
	if ids == "" {
		return c.Status(422).JSON(&fiber.Map{
			"success": false,
			"message": "prodIds params is missing",
			"data":    nil,
		})
	}
	var products []entities.Product
	var idsList []int
	err := json.Unmarshal([]byte(ids), &idsList)
	if err != nil {
		return c.Status(422).JSON(&fiber.Map{
			"success": false,
			"message": "product ids in incorrect formate",
			"data":    nil,
		})
	}
	result := database.Instance.Delete(&products, idsList)
	if result.RowsAffected == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "failed to delete products",
			"data":    nil,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Products deleted successfully",
		"data":    products,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	if c.Locals("user").(auth.ValidUser).Email != "admin" {
		return c.Status(401).JSON(&fiber.Map{
			"success": false,
			"message": "Not authorized for this action",
			"data":    nil,
		})
	}
	if len(c.Body()) == 0 {
		return c.Status(422).JSON(&fiber.Map{
			"success": false,
			"message": "body is missing",
			"data":    nil,
		})
	}
	var product entities.Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(422).JSON(&fiber.Map{
			"success": false,
			"message": "fields are missing",
			"data":    nil,
		})
	}
	result := database.Instance.Updates(&product)
	if result.RowsAffected == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "failed to update product",
			"data":    nil,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Product has been updated successfuly",
		"data":    product,
	})
}
