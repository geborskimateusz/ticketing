package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Called /api/users/signin"})
}
