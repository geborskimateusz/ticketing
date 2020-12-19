package controllers

import (
	"errors"
	"net/http"

	"github.com/geborskimateusz/auth/server/db"
	"github.com/geborskimateusz/auth/server/entity"
	"github.com/geborskimateusz/auth/server/validation"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

func Signup(c *gin.Context) {

	var user entity.User
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err == nil {
		validate := validator.New()
		if err := validate.Struct(&user); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			c.Error(validation.NewRequestValidationError(validationErrors))
			return
		}
	}

	usersFound, err := db.FindBy(bson.M{"email": user.Email})
	if err != nil {
		c.Error(err)
		return
	}

	if len(usersFound) != 0 {
		c.Error(errors.New("Email already in use"))
		return
	}

	_, err = db.Create(user)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, user)

}
