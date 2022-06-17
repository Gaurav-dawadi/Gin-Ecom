package middlewares

import (
	"fmt"
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
		token := strings.Split(header, " ")[1]

		if header == "" {
			fmt.Println(utils.AUTHORIZATION)
			c.JSON(http.StatusUnauthorized, utils.UNAUTHORIZED)
			c.Abort()
			return
		}

		parsed_token, err := jwt.ParseWithClaims(
			token,
			&utils.TokenClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return utils.SignKey, nil
			},
		)

		expiration_time := parsed_token.Claims.(*utils.TokenClaims).ExpiresAt.Time
		if expiration_time.Sub(time.Now().Local()) < 0 {
			fmt.Println(utils.TOKEN_EXPIRED)
			c.JSON(http.StatusUnauthorized, utils.TOKEN_EXPIRED)
			c.Abort()
			return
		}

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				fmt.Println(utils.UNAUTHORIZED)
				c.JSON(http.StatusUnauthorized, utils.UNAUTHORIZED)
				c.Abort()
				return
			}
			if err == jwt.ErrTokenExpired {
				fmt.Println(utils.TOKEN_EXPIRED)
				c.JSON(http.StatusUnauthorized, utils.TOKEN_EXPIRED)
				c.Abort()
				return
			}
			fmt.Println(utils.BAD_REQUEST)
			c.Abort()
			return
		}

		if !parsed_token.Valid {
			fmt.Println(utils.UNAUTHORIZED)
			c.JSON(http.StatusUnauthorized, utils.UNAUTHORIZED)
			c.Abort()
			return
		}

		c.Next()
	}
}

func JWTAuthRefreshToken() {}
