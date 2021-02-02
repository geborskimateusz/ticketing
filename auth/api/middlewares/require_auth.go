package middlewares

import (
	"github.com/geborskimateusz/auth/api/entity"
	"github.com/geborskimateusz/auth/api/validation"
	"github.com/gin-gonic/gin"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Keys["currentUser"].(*entity.UserPayload); !ok {
			c.Error(validation.NewNotAuthorizedError())
			c.Abort()
		}

		c.Next()
	}
}
