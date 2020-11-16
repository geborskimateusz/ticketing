package controllers

import (
	"fmt"
	"net/http"

	"github.com/geborskimateusz/auth/server/entity"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var body entity.User
	err := c.Bind(body)

	if err != nil {
		fmt.Println("afasa")
		// retus error
	}

	c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("%v", body)})
}
