package controllers

import (
	"go-practice/apis/services"
	"go-practice/models"
	"go-practice/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProduct(c *gin.Context) {
	result, err := services.ServiceAllProduct()

	if err != nil {
		res := response.ResponseBadRequest("Failed to Find products")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		res := response.ResponseBadRequest("Some field is not correct in product model")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := services.CreateProduct(product); err != nil {
		res := response.ResponseBadRequest("Couldnot create Product")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusCreated, "Product Created Successfully")
}
