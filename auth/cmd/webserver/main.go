package main

import (
	"log"
	"net/http"

	"github.com/geborskimateusz/auth"
)

func main() {

	server, err := auth.NewAuthServer()

	if err != nil {
		log.Fatalf("problem creating auth http server %v", err)
	}

	if err := http.ListenAndServe(":3000", server); err != nil {
		log.Fatalf("could not listen on port 3000 %v", err)
	}

}
