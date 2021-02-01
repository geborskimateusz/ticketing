package middlewares

import (
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
			c.Next()
			return
		}

		token, _ := jwt.ParseWithClaims(
			cookie.Value,
			&util.CustomClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_KEY")), nil
			},
		)

		claims, _ := token.Claims.(*util.CustomClaims)

		c.Set("currentUser", &UserPayload{ID: claims.ID, Email: claims.Email})

		c.Next()
	}
}
