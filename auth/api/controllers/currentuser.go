package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {
	cookie, err := c.Request.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{currentUser: nil})
		return
	}

	//verify token here https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac

	c.JSON(200, cookie)

}
