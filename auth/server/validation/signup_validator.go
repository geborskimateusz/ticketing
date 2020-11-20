package validation

import (
	"fmt"
	"net/http"

	"github.com/geborskimateusz/auth/server/entity"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// SignupValidator validates request body for /api/users/signup
func SignupValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entity.User
		if err := c.ShouldBindBodyWith(&user, binding.JSON); err == nil {
			validate := validator.New()
			if err := validate.Struct(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": fmt.Sprint(err.Error()),
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
