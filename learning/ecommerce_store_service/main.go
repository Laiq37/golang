package main

import (
	"ecommerce-store-service/auth"
	"ecommerce-store-service/database"
	"ecommerce-store-service/routes"
	"fmt"
)

func main() {
	LoadConfig()

	auth.JwtKey = []byte(AppConfig.SecretKey)

	//initializing Fiber Router/App instance
	routes.InitRouter()

	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	//defining routes
	routes.LoadRoutes()

	routes.App.Listen(fmt.Sprintf(":%d", AppConfig.Port))
}
