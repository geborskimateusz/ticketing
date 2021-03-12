package controllers

import (
	"net/http"

	db "github.com/geborskimateusz/auth/api/db"
	entity "github.com/geborskimateusz/auth/api/entity"
	util "github.com/geborskimateusz/auth/api/util"
	common "github.com/geborskimateusz/ticketing-common"
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
		c.Error(common.NewBadRequestError("Email already in use"))
		return
	}

	saved, err := db.Create(user)
	if err != nil {
		c.Error(err)
		return
	}

	jwtToken, err := util.CreateToken(saved.ID.Hex(), saved.Email)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	util.SetCookie(c, "jwt", jwtToken)

	c.JSON(http.StatusCreated, saved.AsJSON())
}
