package controllers

import (
	"go-practice/apis/repository"
	"go-practice/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProductImage(c *gin.Context) {
	result, err := repository.GetAllProductImage()

	if err != nil {
		res := response.ResponseBadRequest("Failed to Find products")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, result)
}
