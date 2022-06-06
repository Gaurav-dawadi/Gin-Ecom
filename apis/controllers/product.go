package controllers

import (
	"go-practice/apis/services"
	"go-practice/models"
	"go-practice/response"
	"go-practice/utils"
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
	file, uploadFile, err := c.Request.FormFile("Image")

	if err != nil {
		response.ResponseBadRequest("Failed to get file form request")
		return
	}

	// Upload file to a local folder and return it's filepath
	filepath := utils.FileSystemStorage(file, uploadFile)

	var product models.Product

	product_result, err := services.CreateProduct(product)

	if err != nil {
		res := response.ResponseBadRequest("Couldnot create Product")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := services.CreateProductImage(product_result.ID, filepath); err != nil {
		res := response.ResponseBadRequest("Couldnot create ProductImage")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": product_result})
}
