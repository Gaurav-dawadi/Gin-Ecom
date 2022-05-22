package controller

import (
	"go-practice/models"
	"go-practice/response"
	"go-practice/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func GetAllUsers(c *gin.Context) {
	var users []models.User
	result := common.SetupDatabase()
	err := result.Find(&users).Error

	if err != nil {
		res := response.ResponseBadRequest("Failed to Find users")
		c.JSON(http.StatusCreated, res)
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context){
	var user models.User
	user_id, err := strconv.ParseInt(c.Param("id"),10,64)
	
	if err != nil{
		res := response.ResponseBadRequest("User id not provided")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// res := common.SetupDatabase().Model(&models.User{}).Where("id = ?", user_id).First(&user).Error
	res := common.SetupDatabase().First(&user, user_id).Error

	if res != nil {
		res := response.ResponseBadRequest("Failed to Get required users")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context){
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		res := response.ResponseBadRequest("Error in provided data validation")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := common.SetupDatabase().Create(&user).Error;err != nil {
		res := response.ResponseBadRequest("Failed to Create users")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusCreated, user)
}

