package controllers

import (
	b64 "encoding/base64"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {
	cookie, err := c.Request.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"currentUser": nil})
		return
	}

	tokenBytes, err := b64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"currentUser": nil})
		return
	}
	tokenString := string(tokenBytes)

	// verify token here https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
	c.JSON(200, tokenString)

}
