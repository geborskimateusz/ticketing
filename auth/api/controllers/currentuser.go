package controllers

import (
	"net/http"

	common "github.com/geborskimateusz/ticketing-common"
	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {
	currentUser, ok := c.Keys["currentUser"].(*common.UserPayload)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"currentUser": nil})
		return

	}

	c.JSON(200, gin.H{"currentUser": currentUser})
}
