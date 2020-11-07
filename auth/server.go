package auth

import "net/http"

type AuthServer struct {
	http.Handler
}

func NewAuthServer() (*AuthServer, error) {
	a := new(AuthServer)

	router := http.NewServeMux()

	a.Handler = router

	return a, nil
}
