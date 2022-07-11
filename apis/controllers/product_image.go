package controllers

import (
	"go-practice/apis/repository"
	"go-practice/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductImageController struct {
	productImageRepository repository.ProductImageRepository
}

func NewProductImageController(productImageRepository repository.ProductImageRepository) ProductImageController {
	return ProductImageController{
		productImageRepository: productImageRepository,
	}
}

func (pic ProductImageController) GetAllProductImage(c *gin.Context) {
	result, err := pic.productImageRepository.GetAllProductImage()

	if err != nil {
		res := response.ResponseBadRequest("Failed to Find products")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, result)
}
