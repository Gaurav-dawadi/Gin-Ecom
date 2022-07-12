package routes

import (
	"go-practice/apis/controllers"
	"go-practice/apis/middlewares"

	"github.com/gin-gonic/gin"
)

type RouteInitializer struct {
	categoryController     controllers.CategoryController
	commentController      controllers.CommentController
	productController      controllers.ProductController
	userController         controllers.UserController
	productImageController controllers.ProductImageController
}

func NewRouteInitializer(
	categoryController controllers.CategoryController,
	commentController controllers.CommentController,
	productController controllers.ProductController,
	userController controllers.UserController,
	productImageController controllers.ProductImageController,
) *RouteInitializer {
	return &RouteInitializer{
		categoryController:     categoryController,
		commentController:      commentController,
		productController:      productController,
		userController:         userController,
		productImageController: productImageController,
	}
}

func (ri RouteInitializer) RouteSetup() *gin.Engine {
	r := gin.Default()
	router := r.Group("/api")

	router.POST("login", controllers.Login)
	router.POST("signup", ri.userController.CreateUser)
	router.POST("jwt-refresh", controllers.RefreshToken)
	router.POST("jwt-clear", controllers.ClearToken)
	router.Use(middlewares.JwtAuthValidation())

	{
		router.GET("users", ri.userController.GetAllUsers)
		router.GET("user", ri.userController.GetUser)
		router.GET("user/:id", ri.userController.GetUser)
		router.GET("categories", ri.categoryController.GetAllCategory)
		router.POST("category", ri.categoryController.CreateCategory)
		router.GET("products", ri.productController.GetAllProduct)
		router.GET("product", ri.productController.GetProduct)
		router.POST("product", ri.productController.CreateProduct)
		router.PATCH("product/:id", ri.productController.UpdateProduct)
		router.DELETE("product/:id", ri.productController.DeleteProduct)
		router.GET("comments", ri.commentController.GetAllComment)
		router.POST("comment", ri.commentController.CreateComment)
		router.GET("product-image", ri.productImageController.GetAllProductImage)
	}
	return r
}
