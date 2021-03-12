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
		c.Error(common.NewBadRequestError("Invalid Credentials"))
		return
	}

	jwtToken, err := util.CreateToken(userFound.ID.Hex(), userFound.Email)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	util.SetCookie(c, "jwt", jwtToken)

	c.JSON(http.StatusOK, userFound.AsJSON())
}
