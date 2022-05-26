package controllers

import (
	"go-practice/apis/services"
	"go-practice/models"
	"go-practice/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(c *gin.Context) {
	result, err := services.GetAllCategory()
	if err != nil {
		res := response.ResponseBadRequest("Error while Fetching Categories")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		res := response.ResponseBadRequest("Some field is not correct in category model")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := services.CreateCategory(category); err != nil {
		res := response.ResponseBadRequest("Couldnot create Category")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusCreated, category)
}
