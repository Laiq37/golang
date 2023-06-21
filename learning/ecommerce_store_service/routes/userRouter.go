package routes

import "ecommerce-store-service/handlers"

func loadUserRoutes() {
	App.Post("/signup", handlers.Signup)
	App.Post("/login", handlers.Login)
}
