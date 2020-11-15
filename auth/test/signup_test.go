package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/geborskimateusz/auth/server"
	"github.com/geborskimateusz/auth/server/controllers"

	gin "github.com/gin-gonic/gin"
)

func TestSignup(t *testing.T) {

	msg, _ := json.Marshal(map[string]string{
		"email":    "test@gmail.com",
		"password": "123dsfsdf",
	})
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	r.POST(server.SignupRoute, controllers.Signup)
	c.Request, _ = http.NewRequest("POST", server.SignupRoute, bytes.NewBuffer(msg))
	r.ServeHTTP(w, c.Request)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	readBuf, _ := ioutil.ReadAll(w.Body)

	if !reflect.DeepEqual(msg, readBuf) {
		t.Errorf("Expected body %d, got %d", msg, readBuf)
	}
}
