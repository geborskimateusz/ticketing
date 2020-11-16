package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/geborskimateusz/auth/server/entity"

	"github.com/geborskimateusz/auth/server"
	"github.com/geborskimateusz/auth/server/controllers"

	gin "github.com/gin-gonic/gin"
)

func TestSignup(t *testing.T) {

	newUser := entity.User{
		"test@test.com",
		"1234asvvsd",
	}

	userBytes, _ := json.Marshal(newUser)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.POST(server.SignupRoute, controllers.Signup)
	c.Request, _ = http.NewRequest("POST", server.SignupRoute, bytes.NewReader(userBytes))
	r.ServeHTTP(w, c.Request)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	AssertResponseBody(t, w.Body.String(), string(userBytes))
}
