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

	filepath := utils.FileSystemStorage(file, uploadFile)

	var product models.Product

	if err := c.ShouldBind(&product); err != nil {
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

// Todo: Use ProductWithImage model(struct) and take product + image from user.
// Then first create product and use its id for creating productImage model

// func CreateProduct(c *gin.Context) {

// file, uploadFile, err := c.Request.FormFile("file")
// var productImg models.ProductImage
// var productImg models.ProductWithImage

// fmt.Println("Bind: ", file, uploadFile, err)

// fmt.Println(c.Request.FormFile("image"))

// if err := c.ShouldBind(&productImg); err != nil {
// 	res := response.ResponseBadRequest("Some field is not correct in product model")
// 	c.JSON(http.StatusBadRequest, res)
// 	return
// }
// file, uploadFile, err := c.Request.FormFile("file")

// fmt.Println(productImg)

// for _, v := range productImg.Images {
// 	fmt.Println("Inside loop")
// 	f, err := c.FormFile(v)
// 	fmt.Println(f, err)
// }

// c.Bind(&productImg.Images)
// if err := c.Bind(&productImg.Product); err != nil {
// 	res := response.ResponseBadRequest("Some field is not correct in product model")
// 	c.JSON(http.StatusBadRequest, res)
// 	return
// }
// if err := c.Bind(&productImg.Images); err != nil {
// 	res := response.ResponseBadRequest("Some field is not correct in product model")
// 	c.JSON(http.StatusBadRequest, res)
// 	return
// }

// if err := services.CreateProduct(productImg); err != nil {
// 	res := response.ResponseBadRequest("Couldnot create Product")
// 	c.JSON(http.StatusBadRequest, res)
// 	return
// }
// c.JSON(http.StatusCreated, "Product Created Successfully")
// }
