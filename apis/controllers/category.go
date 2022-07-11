package controllers

import (
	"go-practice/apis/services"
	"go-practice/models"
	"go-practice/response"
	"net/http"

	"go-practice/utils/logger"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
	}
}

func (cc CategoryController) GetAllCategory(c *gin.Context) {
	result, err := cc.categoryService.GetAllCategory()
	if err != nil {
		res := response.ResponseBadRequest("Error while Fetching Categories")
		logger.FailOnError(err, "Error while Fetching Categories")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (cc CategoryController) CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		res := response.ResponseBadRequest("Some field is not correct in category model")
		logger.FailOnError(err, "Some field is not correct in category model")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := cc.categoryService.CreateCategory(category); err != nil {
		res := response.ResponseBadRequest("Couldnot create Category")
		logger.FailOnError(err, "Couldnot create Category")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusCreated, category)
}
