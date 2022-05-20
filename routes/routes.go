package routes

import (
	"go-practice/apis/controller"
	"github.com/gin-gonic/gin"
)

func routeSetup() *gin.Engine {
	r := gin.Default()
	router := r.Group("/api")
	{
		router.GET("user", controller.get_all_user())
		router.POST("user", controller.post_user())
		router.GET("user/:id", controller.get_user())
	}
	return r
}
