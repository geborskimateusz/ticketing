package main

import (
	"os"

	"github.com/geborskimateusz/auth/api"
)

func main() {

	if os.Getenv("JWT_KEY") == "" {
		panic("JWT_KEY must be definied")
	}

	api.Instance().Run()

}
