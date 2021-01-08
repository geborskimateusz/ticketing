package main

import "github.com/geborskimateusz/auth/api"

func main() {
	// api.Instance().Run()
	router := api.Instance()
	router.RunTLS(":8080", "./certs/server.crt", "./certs/server.key")
}
