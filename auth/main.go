package auth

import (
	"github.com/geborskimateusz/auth/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	SetupServer().Run()
}

const (
	SignupRoute      string = "/api/users/signup"
	SigninRoute      string = "/api/users/signin"
	SignoutRoute     string = "/api/users/signout"
	CurrentUserRoute string = "/api/users/currentuser"
)

func SetupServer() *gin.Engine {
	router := gin.Default()

	router.POST(CurrentUserRoute, controllers.Signup)
	router.POST(SigninRoute, controllers.Signin)
	router.POST(SignoutRoute, controllers.Signout)
	router.GET(CurrentUserRoute, controllers.CurrentUser)

	return router
}
