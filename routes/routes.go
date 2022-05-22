package routes

import (
	"go-practice/apis/controller"
	"github.com/gin-gonic/gin"
)

func RouteSetup() *gin.Engine {
	r := gin.Default()
	router := r.Group("/api")
	{
		router.GET("user", controller.GetAllUsers)
		router.POST("user", controller.CreateUser)
		router.GET("user/:id", controller.GetUser)
	}
	return r
}
