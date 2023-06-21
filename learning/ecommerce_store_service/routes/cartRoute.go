package routes

import "ecommerce-store-service/handlers"

func loadCartRoutes() {
	SecureApp.Get("/userCart", handlers.GetCartItems)
	SecureApp.Post("/userCart", handlers.AddItemToCart)
	SecureApp.Delete("/userCart/:id", handlers.DeleteCartItem)
	SecureApp.Put("/userCart", handlers.UpdateCart)
}
