package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		// retus error
	}

	fmt.Printf("%s", string(jsonData))
	c.JSON(http.StatusOK, gin.H{"data": jsonData})
}
