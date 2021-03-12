package controllers

import (
	"net/http"

	util "github.com/geborskimateusz/auth/api/util"
	"github.com/gin-gonic/gin"
)

func Signout(c *gin.Context) {
	util.ClearSession(c)
	c.JSON(http.StatusOK, gin.H{})
}
