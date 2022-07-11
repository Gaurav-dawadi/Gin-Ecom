package routes

import (
	"go-practice/apis/controllers"
	"go-practice/apis/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userController controllers.UserController
}

func NewUserRoutes(userController controllers.UserController) *UserRoutes {
	return &UserRoutes{
		userController: userController,
	}
}

func (ur UserRoutes) UserRouteSetup() *gin.Engine {
	r := gin.Default()
	router := r.Group("/api")
	{
		router.POST("signup", ur.userController.CreateUser)
		router.GET("users", middlewares.JwtAuthValidation(), ur.userController.GetAllUsers)
		router.GET("user/:id", middlewares.JwtAuthValidation(), ur.userController.GetUser)
	}
	return r
}
