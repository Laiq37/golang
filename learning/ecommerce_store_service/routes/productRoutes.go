package routes

import "ecommerce-store-service/handlers"

func loadProductRoutes() {
	App.Get("/products", handlers.GetProducts)
	App.Get("/products/:id", handlers.GetProduct)
	SecureApp.Post("/products", handlers.CreateProduct)
	SecureApp.Put("/products", handlers.UpdateProduct)
	SecureApp.Delete("/products/:id", handlers.DeleteProduct)
	SecureApp.Delete("/products", handlers.DeleteProducts)
}
