package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {
	session := sessions.Default(c)
	c.JSON(200, gin.H{"token": session.Get("jwt")})

}
