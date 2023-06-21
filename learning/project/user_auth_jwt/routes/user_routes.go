package routes

import (
	controller "user-auth-jwt/controllers"
	"user-auth-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	//to authenticate use before sending to particular routes
	incomingRoutes.Use(middleware.Authenticate())

	//routes
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/user/:id", controller.GetUser())
}
