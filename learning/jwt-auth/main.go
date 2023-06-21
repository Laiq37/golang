package main

import (
	"jwt-authentication/auth"
	"jwt-authentication/controllers"
	"jwt-authentication/database"
	"jwt-authentication/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	//load Configuration
	LoadAppConfig()

	auth.JwtKey = []byte(AppConfig.SecretKey)

	//Initialize Database
	database.Connect(AppConfig.ConnectString)
	database.Migrate()

	//Initialize Router
	router := initRouter()
	router.Run(":8000")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}

//https://codewithmukesh.com/blog/jwt-authentication-in-golang/
//https://codewithmukesh.com/blog/implementing-crud-in-golang-rest-api/#Connecting_to_the_database
