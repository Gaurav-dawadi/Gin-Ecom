package routes

import (
	"go-practice/apis/controllers"
	"go-practice/apis/middlewares"

	"github.com/gin-gonic/gin"
)

type CategoryRoutes struct {
	categoryController controllers.CategoryController
}

func NewCategoryRoutes(categoryController controllers.CategoryController) *CategoryRoutes {
	return &CategoryRoutes{
		categoryController: categoryController,
	}
}

func (cr CategoryRoutes) CategoryRouteSetup() *gin.Engine {
	r := gin.Default()
	router := r.Group("/api")
	{
		router.GET("categories", middlewares.JwtAuthValidation(), cr.categoryController.GetAllCategory)
		router.POST("category", middlewares.JwtAuthValidation(), cr.categoryController.CreateCategory)
	}
	return r
}
