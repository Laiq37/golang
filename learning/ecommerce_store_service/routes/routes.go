package routes

import (
	"ecommerce-store-service/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
)

var App *fiber.App
var SecureApp fiber.Router

// var adminApp fiber.Router

// middleware := func (c *fiber.CTx) error {
// 	return c.Next()
// }

func InitRouter() {
	log.Println("Initializing Routes ...")
	App = fiber.New()
	SecureApp = App.Group("/api", middleware.VerifyUser)
	log.Println("Router has been Initialized!")
}

func LoadRoutes() {
	log.Println("Loading All routes ...")
	loadUserRoutes()
	loadProductRoutes()
	loadOrderRoutes()
	loadCartRoutes()
	loadPaymentRoutes()
	log.Println("All Routes has been Loaded!")
}
