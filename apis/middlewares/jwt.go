package middlewares

import (
	"fmt"
	"go-practice/response"
	"go-practice/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuthValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(utils.AUTHORIZATION)
		if header == "" {
			fmt.Println(utils.AUTHORIZATION)
			c.JSON(http.StatusUnauthorized, utils.UNAUTHORIZED)
			c.Abort()
			return
		}

		token := strings.Split(header, " ")[1]
		parsed_token, err := JWTParseWithClaims(token)

		expiration_time := parsed_token.Claims.(*utils.TokenClaims).ExpiresAt.Time
		if expiration_time.Sub(time.Now().Local()) < 0 {
			fmt.Println(utils.TOKEN_EXPIRED)
			c.JSON(http.StatusUnauthorized, utils.TOKEN_EXPIRED)
			c.Abort()
			return
		}

		if err != nil {
			res := response.JWTErrorResponse(err)
			c.JSON(res.Status, res.Code)
			c.Abort()
			return
		}

		if !parsed_token.Valid {
			fmt.Println(utils.BAD_REQUEST)
			c.JSON(http.StatusBadRequest, utils.BAD_REQUEST)
			c.Abort()
			return
		}

		c.Set(utils.USER_ID, parsed_token.Claims.(*utils.TokenClaims).UserID)
		c.Set(utils.USER_NAME, parsed_token.Claims.(*utils.TokenClaims).Username)
		c.Next()
	}
}

func JWTParseWithClaims(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(
		token,
		&utils.TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return utils.SignKey, nil
		},
	)
}
