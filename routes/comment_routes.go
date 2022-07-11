package routes

import (
	"go-practice/apis/controllers"
	"go-practice/apis/middlewares"

	"github.com/gin-gonic/gin"
)

type CommentRoutes struct {
	commentController controllers.CommentController
}

func NewCommentRoutes(commentController controllers.CommentController) *CommentRoutes {
	return &CommentRoutes{
		commentController: commentController,
	}
}

func (cr CommentRoutes) CommentRouteSetup() *gin.Engine {
	r := gin.Default()
	router := r.Group("/api")
	{
		router.GET("comments", middlewares.JwtAuthValidation(), cr.commentController.GetAllComment)
		router.POST("comment", middlewares.JwtAuthValidation(), cr.commentController.CreateComment)
	}
	return r
}
