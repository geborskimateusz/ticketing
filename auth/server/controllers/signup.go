package controllers

import (
	"net/http"

	"github.com/geborskimateusz/auth/server/entity"
	"github.com/geborskimateusz/auth/server/errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Signup(c *gin.Context) {

	var user entity.User
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err == nil {
		//throw here request validation error
		validate := validator.New()
		if err := validate.Struct(&user); err != nil {
			c.Error(&errors.RequestValidationError{Err: err.Error()})
			return
		}
	}

	//thro here DatabaseConnectionError
	c.Error(&errors.DatabaseConnectionError{})

	c.JSON(http.StatusOK, user)

}
