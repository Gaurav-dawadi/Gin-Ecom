package controllers

import (
	"fmt"
	"go-practice/apis/services"
	"go-practice/models"
	"go-practice/response"
	"go-practice/utils"
	"go-practice/utils/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService      services.ProductService
	productImageService services.ProductImageService
}

func NewProductController(
	productService services.ProductService,
	productImageService services.ProductImageService,
) *ProductController {
	return &ProductController{
		productService:      productService,
		productImageService: productImageService,
	}
}

func (pc ProductController) GetProduct(c *gin.Context) {
	product_id := c.Param("id")

	product, err := pc.productService.GetProduct(product_id)
	if err != nil {
		logger.LogOutput("Error when fetching product")
		c.JSON(http.StatusBadGateway, "Error when fetching product")
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (pc ProductController) GetAllProduct(c *gin.Context) {
	result, err := pc.productService.GetAllProduct()

	if err != nil {
		res := response.ResponseBadRequest("Failed to Find products")
		logger.FailOnError(err, "Failed to Find products")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (pc ProductController) CreateProduct(c *gin.Context) {
	file, uploadFile, err := c.Request.FormFile("Image")

	if err != nil {
		logger.LogOutput("Failed to get file from request")
		response.ResponseBadRequest("Failed to get file from request")
		return
	}

	// Upload file to a local folder and return it's filepath
	filepath := utils.FileSystemStorage(file, uploadFile)

	var product models.Product

	if err := c.ShouldBind(&product); err != nil {
		res := response.ResponseBadRequest("Couldnot bind Product")
		logger.LogOutput("Couldnot bind Product")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	user_id := c.MustGet(utils.USER_ID)
	int_user_id, ok := utils.TypeAssertInt(user_id)
	if !ok {
		logger.LogOutput("Error during type assertion")
		return
	}

	*product.UserId = uint(int_user_id)
	product_result, err := pc.productService.CreateProduct(product)
	if err != nil {
		res := response.ResponseBadRequest("Couldnot create Product")
		logger.FailOnError(err, "Couldnot create Product")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := pc.productImageService.CreateProductImage(product_result.ID, filepath); err != nil {
		res := response.ResponseBadRequest("Couldnot create ProductImage")
		logger.FailOnError(err, "Couldnot create ProductImage")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": product_result})
}

func (pc ProductController) UpdateProduct(c *gin.Context) {
	product_id := c.Param("id")

	product, err := pc.productService.GetProduct(product_id)
	if err != nil {
		fmt.Println("Error when fetching product")
		c.JSON(http.StatusBadGateway, "Error when fetching product")
		return
	}
	if product.ID == 0 {
		fmt.Println("Error Couldnot find product")
		c.JSON(http.StatusNotFound, "Couldnot find product")
		return
	}

	user_id := c.MustGet(utils.USER_ID)
	typed_user_id, ok := utils.TypeAssertInt(user_id)
	if !ok {
		logger.LogOutput("Error during typed assertion")
		c.JSON(http.StatusInternalServerError, "Error during typed assertion")
		return
	}

	if !utils.IsSameUser(int(*product.UserId), typed_user_id) {
		logger.LogOutput("Error Unauthorized user trying to update product. User and owner of product doesnot match")
		c.JSON(http.StatusBadRequest, "Cannot make this request")
		return
	}

	var product_obj models.UpdateProduct

	if err := c.ShouldBindJSON(&product_obj); err != nil {
		logger.FailOnError(err, "Error binding product")
		c.JSON(http.StatusBadGateway, "Error binding product")
		return
	}

	if product_obj.Name == "" {
		product_obj.Name = product.Name
	}
	if product_obj.CategoryID == nil || *product_obj.CategoryID <= 0 {
		product_obj.CategoryID = product.CategoryID
	}
	if product_obj.Description == "" {
		product_obj.Description = product.Description
	}
	if product_obj.Quantity <= 0 {
		product_obj.Quantity = product.Quantity
	}
	if product_obj.Price <= 0 {
		product_obj.Price = product.Price
	}

	if err := pc.productService.UpdateProduct(product_id, product_obj); err != nil {
		logger.FailOnError(err, "Error updating product")
		c.JSON(http.StatusInternalServerError, "Error updating product")
	}

	c.JSON(http.StatusCreated, "Product Updated Successfully")
}

func (pc ProductController) DeleteProduct(c *gin.Context) {
	product_id := c.Param("id")

	product, err := pc.productService.GetProduct(product_id)
	if err != nil {
		logger.LogOutput("Error when fetching product")
		c.JSON(http.StatusBadGateway, "Error when fetching product")
		return
	}

	user_id := c.MustGet(utils.USER_ID)
	typed_user_id, ok := utils.TypeAssertInt(user_id)
	if !ok {
		logger.LogOutput("Error during typed assertion")
		c.JSON(http.StatusInternalServerError, "Error during typed assertion")
		return
	}

	if !utils.IsSameUser(int(*product.UserId), typed_user_id) {
		logger.LogOutput("Error Unauthorized user trying to delete product.User and owner of product doesnot match")
		c.JSON(http.StatusBadRequest, "Cannot make this request")
		return
	}

	if err := pc.productService.DeleteProduct(product.ID); err != nil {
		logger.LogOutput("Error deleting product")
		c.JSON(http.StatusInternalServerError, "Error deleting product")
	}

	c.JSON(http.StatusOK, "Product Deleted Successfully")
}
