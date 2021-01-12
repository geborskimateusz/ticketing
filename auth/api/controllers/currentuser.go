package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {
	session, _ := Store.Get(c.Request, "cookie-name")
	v := session.Values["jwt"]

	log.Println(" ->>  ", v)
	c.JSON(http.StatusOK, gin.H{"data": "Called /api/users/currentuser"})
}
