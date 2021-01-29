package middlewares

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/geborskimateusz/auth/api/util"
	"github.com/gin-gonic/gin"
)

type UserPayload struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("jwt")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"currentUser": nil})
			return
		}

		token, err := jwt.ParseWithClaims(
			cookie.Value,
			&util.CustomClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_KEY")), nil
			},
		)

		claims, ok := token.Claims.(*util.CustomClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"currentUser": nil})
			return
		}

		c.Next()
	}
}
