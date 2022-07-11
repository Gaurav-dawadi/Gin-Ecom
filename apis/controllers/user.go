package controllers

import (
	"go-practice/apis/services"
	"go-practice/models"
	"go-practice/response"
	"go-practice/utils/logger"
	"net/http"
	"net/mail"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(c *gin.Context) {
	result, err := services.GetAllUsers()

	if err != nil {
		logger.FailOnError(err, "failed to find users")
		res := response.ResponseBadRequest("Failed to Find users")
		c.JSON(http.StatusCreated, res)
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetUser(c *gin.Context) {
	user_id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		logger.FailOnError(err, "user id not provided")
		res := response.ResponseBadRequest("User id not provided")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	user_obj, res := services.GetUser(user_id)

	if res != nil {
		logger.FailOnError(res, "Failed to Get required users")
		res := response.ResponseBadRequest("Failed to Get required users")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	c.JSON(http.StatusOK, user_obj)
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		logger.FailOnError(err, "error during data binding")
		res := response.ResponseBadRequest("Error in provided data validation")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if _, err := mail.ParseAddress(user.Email); err != nil {
		logger.FailOnError(err, "email address is invalid")
		res := response.ResponseBadRequest("Email address is invalid")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	hashedPassword, passwordError := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if passwordError != nil {
		logger.FailOnError(passwordError, "password error")
		res := response.ResponseBadRequest("Password Error")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	user.Password = string(hashedPassword)

	if err := services.CreateUser(user); err != nil {
		logger.FailOnError(err, "failed to create user")
		res := response.ResponseBadRequest("Failed to Create users")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusCreated, "User created Successfully")
}
