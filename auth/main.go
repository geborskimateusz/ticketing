package auth

import (
	"github.com/geborskimateusz/auth/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	SetupServer().Run()
}

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

// SetupServer initilizes http server
func SetupServer() *gin.Engine {
	router := gin.Default()

	router.POST(SignupRoute, controllers.Signup)
	router.POST(SigninRoute, controllers.Signin)
	router.POST(SignoutRoute, controllers.Signout)
	router.GET(CurrentUserRoute, controllers.CurrentUser)

	return router
}
