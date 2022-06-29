package controllers

import (
	"fmt"
	"go-practice/apis/services"
	"go-practice/models"
	"go-practice/response"
	"go-practice/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type updateProduct struct {
	Name        string `json:"name"`
	CategoryID  uint   `json:"category_id"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
}

func GetProduct(c *gin.Context) {
	product_id := c.Param("id")

	product, err := services.GetProduct(product_id)
	if err != nil {
		fmt.Println("Error when fetching product")
		c.JSON(http.StatusBadGateway, "Error when fetching product")
		return
	}
	if product.ID == 0 {
		fmt.Println("Couldnot find product")
		c.JSON(http.StatusNotFound, "Couldnot find product")
		return
	}

	c.JSON(http.StatusCreated, product)
}

func GetAllProduct(c *gin.Context) {
	result, err := services.GetAllProduct()

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
		response.ResponseBadRequest("Failed to get file from request")
		return
	}

	// Upload file to a local folder and return it's filepath
	filepath := utils.FileSystemStorage(file, uploadFile)

	var product models.Product

	if err := c.ShouldBind(&product); err != nil {
		res := response.ResponseBadRequest("Couldnot bind Product")
		c.JSON(http.StatusBadRequest, res)
		return
	}

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

func UpdateProduct(c *gin.Context) {
	product_id := c.Param("id")

	product, err := services.GetProduct(product_id)
	if err != nil {
		fmt.Println("Error when fetching product")
		c.JSON(http.StatusBadGateway, "Error when fetching product")
		return
	}
	if product.ID == 0 {
		fmt.Println("Couldnot find product")
		c.JSON(http.StatusNotFound, "Couldnot find product")
		return
	}

	var product_obj updateProduct

	if err := c.ShouldBind(&product_obj); err != nil {
		fmt.Println("Error binding product")
		c.JSON(http.StatusBadGateway, "Error binding product")
		return
	}

	product_obj_bind := map[string]interface{}{
		"Name":        product_obj.Name,
		"CategoryID":  product_obj.CategoryID,
		"description": product_obj.Description,
		"quantity":    product_obj.Quantity,
		"price":       product_obj.Price,
	}

	if product_obj.Name != "" {
		product_obj_bind["name"] = product_obj.Name
	}
	if product_obj.CategoryID != 0 {
		product_obj_bind["category_id"] = product_obj.CategoryID
	}
	if product_obj.Description != "" {
		product_obj_bind["description"] = product_obj.Description
	}
	if product_obj.Quantity > -1 {
		product_obj_bind["quantity"] = product_obj.Quantity
	}
	if product_obj.Price > -1 {
		product_obj_bind["price"] = product_obj.Price
	}

	updated_product, _ := services.UpdateProduct(product_id, product_obj_bind)

	c.JSON(http.StatusCreated, updated_product)
}
