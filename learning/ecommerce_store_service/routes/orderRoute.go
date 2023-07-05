package routes

import "ecommerce-store-service/handlers"

func loadOrderRoutes() {
	SecureApp.Post("/orders", handlers.CreateOrder)
	SecureApp.Get("/orders", handlers.GetOrders)
}
