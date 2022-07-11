package routes

import (
	"go-practice/apis/controllers"
	"go-practice/apis/middlewares"

	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	productController controllers.ProductController
}

func NewProductRoutes(productController controllers.ProductController) *ProductRoutes {
	return &ProductRoutes{
		productController: productController,
	}
}

func (pr ProductRoutes) ProductRouteSetup() *gin.Engine {
	r := gin.Default()
	router := r.Group("/api")
	{
		router.GET("products", middlewares.JwtAuthValidation(), pr.productController.GetAllProduct)
		router.GET("product", middlewares.JwtAuthValidation(), pr.productController.GetProduct)
		router.POST("product", middlewares.JwtAuthValidation(), pr.productController.CreateProduct)
		router.PATCH("product/:id", middlewares.JwtAuthValidation(), pr.productController.UpdateProduct)
		router.DELETE("product/:id", middlewares.JwtAuthValidation(), pr.productController.DeleteProduct)
	}
	return r
}
