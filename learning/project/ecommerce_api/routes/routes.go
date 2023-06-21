package routes

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.Login())
	incomingRoutes.POST("/users/login", controllers.Signup())
	incomingRoutes.POST("/users/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controller.SearchProduct())
	incomingRoutes.GET("/users/search", controller.SearchProductByQuery())
}
