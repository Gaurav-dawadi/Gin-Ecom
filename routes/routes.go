package routes

import (
	"go-practice/apis/controllers"
	"go-practice/apis/middlewares"

	"github.com/gin-gonic/gin"
)

func RouteSetup() *gin.Engine {
	r := gin.Default()
	router := r.Group("/api")
	{
		router.POST("login", controllers.Login)
		router.POST("signup", controllers.CreateUser)
		router.GET("users", middlewares.JwtAuthValidation(), controllers.GetAllUsers)
		router.GET("user", middlewares.JwtAuthValidation(), controllers.GetUser)
		router.GET("user/:id", middlewares.JwtAuthValidation(), controllers.GetUser)
		router.GET("category", middlewares.JwtAuthValidation(), controllers.GetAllCategory)
		router.POST("category", middlewares.JwtAuthValidation(), controllers.CreateCategory)
		router.GET("product", middlewares.JwtAuthValidation(), controllers.GetAllProduct)
		router.POST("product", middlewares.JwtAuthValidation(), controllers.CreateProduct)
		router.GET("comment", middlewares.JwtAuthValidation(), controllers.GetAllComment)
		router.POST("comment", middlewares.JwtAuthValidation(), controllers.CreateComment)
		router.GET("product-image", middlewares.JwtAuthValidation(), controllers.GetAllProductImage)
	}
	return r
}
