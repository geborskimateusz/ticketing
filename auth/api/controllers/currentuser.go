package controllers

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/geborskimateusz/auth/api/util"
	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {
	cookie, err := c.Request.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"currentUser": nil})
		return
	}

	token, err := jwt.ParseWithClaims(
		cookie.Value,
		&util.CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		},
	)

	claims, ok := token.Claims.(*util.CustomClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"info": "Couldn't parse claims"})
		return
	}

	c.JSON(200, gin.H{"currentUser": claims})
}
