package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/geborskimateusz/auth/server/entity"

	"github.com/geborskimateusz/auth/server"
	"github.com/geborskimateusz/auth/server/controllers"

	gin "github.com/gin-gonic/gin"
)

func TestSignup(t *testing.T) {

	t.Run("should validate and return 200", func(t *testing.T) {
		newUser := entity.User{
			Email:    "test@test.com",
			Password: "1234asvvsd",
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
	})

	t.Run("should fail on validation of any filed then return 400", func(t *testing.T) {
		newUser := entity.User{
			Email:    "testtest.com",
			Password: "1234",
		}

		userBytes, _ := json.Marshal(newUser)
		fmt.Println("prepare")

		w := httptest.NewRecorder()
		c, r := gin.CreateTestContext(w)
		r.POST(server.SignupRoute, controllers.Signup)
		c.Request, _ = http.NewRequest("POST", server.SignupRoute, bytes.NewReader(userBytes))
		r.ServeHTTP(w, c.Request)
		fmt.Println("After")
		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}

		AssertResponseBody(t, w.Body.String(), string(userBytes))
	})
}
