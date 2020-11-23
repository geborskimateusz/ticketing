package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {

	// body := entity.User{}

	// if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
	// 	c.AbortWithStatusJSON(
	// 		http.StatusInternalServerError,
	// 		gin.H{"error": err.Error()})
	// 	return
	// }
	// fmt.Println(body)

	// c.JSON(http.StatusOK, body)

	c.Error(errors.New("Something went wrong during signup"))
}
