package utils

import (
	"go-practice/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var SignKey = []byte(SECRET_KEY)

type TokenManager struct {
	AcessToken   string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenClaims struct {
	*jwt.RegisteredClaims
	TokenType string `json:"token_type"`
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
}

func create_access_token(user models.User) string {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = &TokenClaims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ACCESS_TOKEN_EXPIRY_TIME)),
		},
		"access",
		int(user.ID),
		user.Username,
	}
	result, _ := token.SignedString(SignKey)
	return result
}

func create_refresh_token(user models.User) string {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = &TokenClaims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(REFRESH_TOKEN_EXPIRY_TIME)),
		},
		"refresh",
		int(user.ID),
		user.Username,
	}
	result, _ := token.SignedString(SignKey)
	return result
}

func JwtAuthToken(user models.User) TokenManager {
	access_token := create_access_token(user)
	refresh_token := create_refresh_token(user)

	var token_manager TokenManager
	token_manager.AcessToken = access_token
	token_manager.RefreshToken = refresh_token

	return token_manager
}
