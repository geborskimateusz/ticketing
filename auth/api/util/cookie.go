package util

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetCookie(c *gin.Context, name, value string) {
	expire := time.Now().Add(20 * time.Minute) // Expires in 20 minutes
	cookie := http.Cookie{Name: name, Value: value, Path: "/", Expires: expire, MaxAge: 86400, HttpOnly: true, Secure: true}
	http.SetCookie(c.Writer, &cookie)
}

func ClearSession(c *gin.Context) {
	cookie := http.Cookie{
		Name:   "jwt",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(c.Writer, &cookie)
}
