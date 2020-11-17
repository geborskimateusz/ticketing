package controllers

import (
	"net/http"

	"github.com/geborskimateusz/auth/server/entity"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {

	body := entity.User{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, body)
}
