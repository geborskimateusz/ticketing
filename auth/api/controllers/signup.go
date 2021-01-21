package controllers

import (
	"net/http"

	"github.com/geborskimateusz/auth/api/db"
	"github.com/geborskimateusz/auth/api/entity"
	"github.com/geborskimateusz/auth/api/util"
	"github.com/geborskimateusz/auth/api/validation"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Signup creates new user
func Signup(c *gin.Context) {
	var user entity.User
	c.ShouldBindBodyWith(&user, binding.JSON)

	userFound, err := db.FindBy(db.Filter("email", user.Email))
	if err != nil {
		c.Error(err)
		return
	}

	if userFound != (entity.UserDoc{}) {
		c.Error(validation.NewBadRequestError("Email already in use"))
		return
	}

	saved, err := db.Create(user)
	if err != nil {
		c.Error(err)
		return
	}

	encodedToken, err := util.GenerateJWTtoken(saved.ID.String(), saved.Email)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	util.SetCookie(c, "jwt", encodedToken)

	c.JSON(http.StatusCreated, saved.AsJSON())
}
