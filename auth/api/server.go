package api

import (
	controllers "github.com/geborskimateusz/auth/api/controllers"
	middlewares "github.com/geborskimateusz/auth/api/middlewares"
	common "github.com/geborskimateusz/ticketing-common"
	"github.com/gin-gonic/gin"
)

const (
	// SignupRoute is used to create new user
	SignupRoute string = "/api/users/signup"

	// SigninRoute is used to log in as existing user
	SigninRoute string = "/api/users/signin"

	// SignoutRoute is used to log out current user
	SignoutRoute string = "/api/users/signout"

	// CurrentUserRoute obtain current user
	CurrentUserRoute string = "/api/users/currentuser"
)

// Instance initilizes http server
func Instance() *gin.Engine {

	router := gin.Default()

	router.Use(common.ErrorHandler())

	router.POST(SignupRoute, middlewares.ValidateRequest(), controllers.Signup)
	router.POST(SigninRoute, middlewares.ValidateRequest(), controllers.Signin)
	router.POST(SignoutRoute, controllers.Signout)
	router.GET(CurrentUserRoute, common.CurrentUser(), controllers.CurrentUser)

	router.NoRoute(func(c *gin.Context) {
		c.Error(common.NewNotFoundError())
	})

	return router
}
