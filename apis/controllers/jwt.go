package controllers

import (
	"go-practice/apis/services"
	"go-practice/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context) {
	var user_info UserCredentials

	if err := ctx.ShouldBindJSON(&user_info); err != nil {
		ctx.JSON(http.StatusBadRequest, "error found in given data")
		return
	}

	user_obj, err := services.GetUserFromEmail(user_info.Email)
	if err != nil {
		// s := fmt.Errorf("error while fetching user from given email: %s", err)  /// Doesnot work
		// s := fmt.Sprintf("error while fetching user from given email: %s", err) /// Works
		ctx.JSON(http.StatusBadRequest, "error while fetching user from given email")
		return
	}
	if user_obj.ID == 0 {
		ctx.JSON(http.StatusBadRequest, "incorrect email address")
		return
	}

	password_db := []byte(user_obj.Password)
	password_client := []byte(user_info.Password)

	if err := bcrypt.CompareHashAndPassword(password_db, password_client); err != nil {
		ctx.JSON(http.StatusBadRequest, "incorrect password")
		return
	}

	token_result := utils.JwtAuthToken(*user_obj)
	token_obj := map[string]interface{}{
		"access_token":  token_result.AcessToken,
		"refresh_token": token_result.RefreshToken,
		"user_detail": map[string]interface{}{
			"id":       user_obj.ID,
			"username": user_obj.Username,
		},
	}

	ctx.JSON(http.StatusOK, token_obj)
}
