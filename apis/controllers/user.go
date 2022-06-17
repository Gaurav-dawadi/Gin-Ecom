package controllers

import (
	"go-practice/apis/services"
	"go-practice/models"
	"go-practice/response"
	"net/http"
	"net/mail"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(c *gin.Context) {
	result, err := services.GetAllUsers()

	if err != nil {
		res := response.ResponseBadRequest("Failed to Find users")
		c.JSON(http.StatusCreated, res)
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetUser(c *gin.Context) {
	user_id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		res := response.ResponseBadRequest("User id not provided")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	user_obj, res := services.GetUser(user_id)

	if res != nil {
		res := response.ResponseBadRequest("Failed to Get required users")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	c.JSON(http.StatusOK, user_obj)
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		res := response.ResponseBadRequest("Error in provided data validation")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if _, err := mail.ParseAddress(user.Email); err != nil {
		res := response.ResponseBadRequest("Email address is invalid")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	hashedPassword, passwordError := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if passwordError != nil {
		res := response.ResponseBadRequest("Password Error")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	user.Password = string(hashedPassword)

	if err := services.CreateUser(user); err != nil {
		res := response.ResponseBadRequest("Failed to Create users")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusCreated, "User created Successfully")
}
