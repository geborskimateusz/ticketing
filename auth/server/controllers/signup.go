package controllers

import (
	"fmt"

	"github.com/geborskimateusz/auth/server/customerr"
	"github.com/geborskimateusz/auth/server/entity"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Signup(c *gin.Context) {

	fmt.Println("In controller")
	var user entity.User
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err == nil {
		//throw here request validation error
		validate := validator.New()
		if err := validate.Struct(&user); err != nil {
			c.Error(&customerr.RequestValidationError{Errors: err})
			return
		}
	}

	//thro here DatabaseConnectionError
	c.Error(&customerr.DatabaseConnectionError{})

	// c.JSON(http.StatusOK, user)

}
