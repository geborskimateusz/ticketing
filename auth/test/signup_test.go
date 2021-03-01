package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/geborskimateusz/auth/api"
)

func TestSignupRoute(t *testing.T) {
	ts := httptest.NewServer(api.Instance())
	defer ts.Close()

	user, _ := json.Marshal(map[string]string{
		"email":    "test@example.com",
		"password": "123123asdsf",
	})
	responseBody := bytes.NewBuffer(user)

	resp, err := http.Post(fmt.Sprintf("%s/api/users/signup", ts.URL), "application/json", responseBody)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	AssertStatusCode(http.StatusCreated, resp.StatusCode)

	val, ok := resp.Header["Content-Type"]

	// Assert that the "content-type" header is actually set
	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	// Assert that it was set as expected
	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}
