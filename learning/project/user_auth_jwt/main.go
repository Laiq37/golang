package main

import (
	"os"
	routes "user-auth-jwt/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//port from env
	port := os.Getenv("PORT")

	//if port empty then assign
	if port == "" {
		port = "8000"
	}

	//creating router
	router := gin.New()

	//setup serverlogger(show logs)
	router.Use(gin.Logger())

	//setting all routes
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	//start server
	router.Run(":" + port)
}
