package auth

import (
	"encoding/json"
	"net/http"
)

// AuthServer is a HTTP interface for authentication.
type AuthServer struct {
	http.Handler
}

const jsonContentType = "application/json"

// NewAuthServer creates a AuthServer with routing configured.
func NewAuthServer() (*AuthServer, error) {
	a := new(AuthServer)

	router := http.NewServeMux()
	router.Handle("/api/users/currentuser", http.HandlerFunc(a.currentUserHandler))

	a.Handler = router

	return a, nil
}

func (a AuthServer) currentUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode("this is a test from server ")
}
