package controllers

import (
	"net/http"

	"github.com/geborskimateusz/auth/api/entity"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Signin(c *gin.Context) {

	var user entity.User
	c.ShouldBindBodyWith(&user, binding.JSON)
	c.JSON(http.StatusOK, user)
}
