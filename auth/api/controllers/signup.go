package controllers

import (
	"log"
	"net/http"
	"time"

	b64 "encoding/base64"

	"github.com/geborskimateusz/auth/api/db"
	"github.com/geborskimateusz/auth/api/entity"
	"github.com/geborskimateusz/auth/api/util"
	"github.com/geborskimateusz/auth/api/validation"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Signup creates new user
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

	usersFound, err := db.FindBy(db.Filter("email", user.Email))
	log.Println(usersFound)
	if err != nil {
		c.Error(err)
		return
	}

	if len(usersFound) != 0 {
		c.Error(validation.NewBadRequestError("Email already in use"))
		return
	}

	saved, err := db.Create(user)
	if err != nil {
		c.Error(err)
		return
	}

	token, err := util.CreateToken(saved.ID.String(), saved.Email)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	encodedToken := b64.StdEncoding.EncodeToString([]byte(token.AccessToken))

	expire := time.Now().Add(20 * time.Minute) // Expires in 20 minutes
	cookie := http.Cookie{Name: "jwt", Value: encodedToken, Path: "/", Expires: expire, MaxAge: 86400, HttpOnly: true, Secure: true}
	http.SetCookie(c.Writer, &cookie)

	c.JSON(http.StatusCreated, saved.AsJSON())
}
