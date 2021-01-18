package controllers

import (
	"net/http"

	"github.com/geborskimateusz/auth/api/entity"
	"github.com/geborskimateusz/auth/api/validation"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Signin(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err == nil {
		validate := validator.New()
		if err := validate.Struct(&user); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			c.Error(validation.NewRequestValidationError(validationErrors))
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": "Called /api/users/signin"})
}
