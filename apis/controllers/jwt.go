package controllers

import (
	"fmt"
	"go-practice/apis/middlewares"
	"go-practice/apis/services"
	"go-practice/response"
	"go-practice/utils"
	"go-practice/utils/logger"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
Checks if user is authentic. If authentic then returns access_token, refresh_token and
some user info
*/
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
		logger.FailOnError(err, "error while fetching user from given email")
		return
	}
	if user_obj.ID == 0 {
		ctx.JSON(http.StatusBadRequest, "incorrect email address")
		logger.LogOutput("incorrect email address")
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

	parsed_token, _ := middlewares.JWTParseWithClaims(token_result.RefreshToken)
	token_time := parsed_token.Claims.(*utils.TokenClaims).ExpiresAt.Time

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "refresh",
		Value:    token_result.RefreshToken,
		Expires:  token_time,
		Secure:   false,
		HttpOnly: true,
	})

	// maxlifetime := math.Max(0, float64(token_time.Day()))
	// ctx.SetCookie("refresh", token_result.RefreshToken, int(maxlifetime), "", "", false, true)

	ctx.JSON(http.StatusOK, token_obj)
}

/*
Takes refresh token either from user as input or search in cookie.
If found returns new access token
*/
func RefreshToken(ctx *gin.Context) {
	var token utils.TokenManager

	if err := ctx.ShouldBindJSON(&token); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if token.RefreshToken == "" {
		cookie, err := ctx.Request.Cookie("refresh")
		if err != nil {
			// Delete cookie refresh code in response
			http.SetCookie(ctx.Writer, &http.Cookie{
				Name:     "refresh",
				Value:    "",
				MaxAge:   -1,
				Secure:   false,
				HttpOnly: true,
			})

			ctx.JSON(http.StatusUnauthorized, "No Refresh Token")
			return
		}
		token.RefreshToken = cookie.Value
	}

	// Validate refreshToken
	parsed_token, err := middlewares.JWTParseWithClaims(token.RefreshToken)
	if err != nil {
		res := response.JWTErrorResponse(err)
		ctx.JSON(res.Status, &gin.H{"Error": res})
		return
	}
	if !parsed_token.Valid {
		logger.LogOutput(utils.BAD_REQUEST)
		ctx.JSON(http.StatusBadRequest, utils.BAD_REQUEST)
		return
	}

	// Check expiry_date of token.RefreshToken
	expiration_time := parsed_token.Claims.(*utils.TokenClaims).ExpiresAt.Time
	if expiration_time.Sub(time.Now().Local()) < 0 {
		logger.LogOutput(utils.TOKEN_EXPIRED)
		ctx.JSON(http.StatusUnauthorized, utils.TOKEN_EXPIRED)
		return
	}

	// If not expired generate new accessToken
	// Else returned refreshToken expired
	decode_user_id := parsed_token.Claims.(*utils.TokenClaims).UserID
	string_user_id := fmt.Sprintf("%d", decode_user_id)
	user_id, err := strconv.ParseInt(string_user_id, 10, 64)
	if err != nil {
		logger.LogOutput("Error while parsing user id to 64bit")
		ctx.JSON(http.StatusInternalServerError, "Error while parsing user id to 64bit")
		return
	}

	user_obj, err := services.GetUser(user_id)
	if err != nil {
		logger.FailOnError(err, "User not found")
		ctx.JSON(http.StatusBadRequest, "User not found")
		return
	}

	new_access_token := utils.CreateAccessToken(*user_obj)
	token_obj := map[string]interface{}{
		"access_token":  new_access_token,
		"refresh_token": token.RefreshToken,
		"user_detail": map[string]interface{}{
			"id":       user_obj.ID,
			"username": user_obj.Username,
		},
	}

	ctx.JSON(http.StatusOK, token_obj)
}

/*
Deletes/Clears refresh token from cookie
*/
func ClearToken(ctx *gin.Context) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "refresh",
		Value:    "",
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
	})

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "access",
		Value:    "",
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
	})

	ctx.JSON(http.StatusOK, "Cookie cleared")
}
