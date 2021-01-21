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

func Signin(c *gin.Context) {

	var user entity.User
	c.ShouldBindBodyWith(&user, binding.JSON)

	userFound, err := db.FindBy(db.Filter("email", user.Email))
	if err != nil {
		c.Error(err)
		return
	}

	passwordMatch := util.ComparePasswords(userFound.Password, user.Password)

	if !passwordMatch {
		c.Error(validation.NewBadRequestError("Invalid Credentials"))
		return
	}

	encodedToken, err := util.GenerateJWTtoken(userFound.ID.String(), userFound.Email)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	util.SetCookie(c, "jwt", encodedToken)

	c.JSON(http.StatusOK, userFound.AsJSON())
}
