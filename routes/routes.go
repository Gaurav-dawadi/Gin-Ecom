package routes

import (
	"go-practice/apis/controllers"

	"github.com/gin-gonic/gin"
)

func RouteSetup() *gin.Engine {
	r := gin.Default()
	router := r.Group("/api")
	{
		router.GET("user", controllers.GetAllUsers)
		router.POST("user", controllers.CreateUser)
		router.GET("user/:id", controllers.GetUser)
		router.GET("category", controllers.GetAllCategory)
		router.POST("category", controllers.CreateCategory)
		router.GET("product", controllers.GetAllProduct)
		router.POST("product", controllers.CreateProduct)
		router.GET("comment", controllers.GetAllComment)
		router.POST("comment", controllers.CreateComment)
		router.GET("product-image", controllers.GetAllProductImage)
	}
	return r
}
