package utils

import (
	"go-practice/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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

func CreateAccessToken(user models.User) string {
	token := jwt.New(jwt.SigningMethodHS256)
	id, _ := uuid.NewRandom()
	token.Claims = &TokenClaims{
		&jwt.RegisteredClaims{
			ID:        id.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ACCESS_TOKEN_EXPIRY_TIME)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
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
	id, _ := uuid.NewRandom()
	token.Claims = &TokenClaims{
		&jwt.RegisteredClaims{
			ID:        id.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(REFRESH_TOKEN_EXPIRY_TIME)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		"refresh",
		int(user.ID),
		user.Username,
	}
	result, _ := token.SignedString(SignKey)
	return result
}

func JwtAuthToken(user models.User) TokenManager {
	access_token := CreateAccessToken(user)
	refresh_token := create_refresh_token(user)

	var token_manager TokenManager
	token_manager.AcessToken = access_token
	token_manager.RefreshToken = refresh_token

	return token_manager
}
