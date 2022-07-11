package controllers

import (
	"go-practice/apis/services"
	"go-practice/models"
	"go-practice/response"
	"go-practice/utils/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllComment(c *gin.Context) {
	result, err := services.GetAllComment()
	if err != nil {
		logger.FailOnError(err, "comments not found")
		res := response.ResponseBadRequest("Comments Not Found")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		logger.FailOnError(err, "some field is not correct in comment")
		res := response.ResponseBadRequest("Some field is not correct in comment")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := services.CreateComment(comment); err != nil {
		logger.FailOnError(err, "coulnot create comment")
		res := response.ResponseBadRequest("Couldnot create comment")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	c.JSON(http.StatusCreated, comment)
}
