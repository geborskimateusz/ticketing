package main

import "github.com/geborskimateusz/auth/api"

func main() {
	// api.Instance().Run()
	router := api.Instance()
	router.RunTLS(":8080", "./fixtures/testdata/cert.pem", "./fixtures/testdata/key.pem")
}
