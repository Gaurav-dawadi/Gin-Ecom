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

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc UserController) GetAllUsers(c *gin.Context) {
	result, err := uc.userService.GetAllUsers()

	if err != nil {
		logger.FailOnError(err, "failed to find users")
		res := response.ResponseBadRequest("Failed to Find users")
		c.JSON(http.StatusCreated, res)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (uc UserController) GetUser(c *gin.Context) {
	user_id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		logger.FailOnError(err, "user id not provided")
		res := response.ResponseBadRequest("User id not provided")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	user_obj, res := uc.userService.GetUser(user_id)

	if res != nil {
		logger.FailOnError(res, "Failed to Get required users")
		res := response.ResponseBadRequest("Failed to Get required users")
		c.JSON(http.StatusBadRequest, res)
		return
	}

	c.JSON(http.StatusOK, user_obj)
}

func (uc UserController) CreateUser(c *gin.Context) {
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

	if err := uc.userService.CreateUser(user); err != nil {
		logger.FailOnError(err, "failed to create user")
		res := response.ResponseBadRequest("Failed to Create users")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	c.JSON(http.StatusCreated, "User created Successfully")
}
