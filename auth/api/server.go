package api

import (
	"github.com/geborskimateusz/auth/api/controllers"
	"github.com/geborskimateusz/auth/api/middlewares"
	"github.com/geborskimateusz/auth/api/validation"
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

	// store := cookie.NewStore([]byte("secret"))
	// store.Options(sessions.Options{Secure: true})

	// router.Use(sessions.Sessions("mysession", store))
	router.Use(middlewares.ErrorHandler())

	router.POST(SignupRoute, controllers.Signup)
	router.POST(SigninRoute, controllers.Signin)
	router.POST(SignoutRoute, controllers.Signout)
	router.GET(CurrentUserRoute, controllers.CurrentUser)

	router.NoRoute(func(c *gin.Context) {
		c.Error(validation.NewNotFoundError())
	})

	return router
}
