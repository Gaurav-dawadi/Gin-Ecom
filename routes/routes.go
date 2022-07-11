package routes

import (
	"go-practice/apis/controllers"

	"github.com/gin-gonic/gin"
)

func RouteSetup() *gin.Engine {
	r := gin.Default()
	router := r.Group("/api")
	{
		router.POST("login", controllers.Login)
		router.POST("jwt-refresh", controllers.RefreshToken)
		router.POST("jwt-clear", controllers.ClearToken)
		// router.GET("product-image", middlewares.JwtAuthValidation(), controllers.GetAllProductImage)
	}
	return r
}
