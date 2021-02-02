package controllers

import (
	"net/http"

	"github.com/geborskimateusz/auth/api/entity"
	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {
	currentUser, ok := c.Keys["currentUser"].(*entity.UserPayload)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"currentUser": nil})
		return

	}

	c.JSON(200, gin.H{"currentUser": currentUser})
}
