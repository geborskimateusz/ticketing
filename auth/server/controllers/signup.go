package controllers

import (
	"github.com/geborskimateusz/auth/server/customerr"
	"github.com/geborskimateusz/auth/server/entity"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Signup(c *gin.Context) {

	var user entity.User
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err == nil {
		validate := validator.New()
		if err := validate.Struct(&user); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			c.Error(customerr.NewRequestValidationError(validationErrors))
			return
		}
	}
	c.Error(customerr.NewDataBaseConnectionError())
}
