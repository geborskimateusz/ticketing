package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/geborskimateusz/auth/server"
)

func TestSignupRoute(t *testing.T) {
	// The setupServer method, that we previously refactored
	// is injected into a test server
	ts := httptest.NewServer(server.SetupServer())
	// Shut down the server and block until all requests have gone through
	defer ts.Close()

	t.Run("should validate api/users/signup", func(t *testing.T) {

		// Make a request to our server with the {base url}/ping
		requestBody, err := json.Marshal(map[string]string{
			"email":    "test@gmail.com",
			"password": "123dsfsdf",
		})
		if err != nil {
			print(err)
		}
		resp, err := http.Post(fmt.Sprintf("%s"+server.SignupRoute, ts.URL), "application/json", bytes.NewBuffer(requestBody))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.StatusCode != 200 {
			t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
		}

		val, ok := resp.Header["Content-Type"]

		// Assert that the "content-type" header is actually set
		if !ok {
			t.Fatalf("Expected Content-Type header to be set")
		}

		// Assert that it was set as expected
		if val[0] != "application/json; charset=utf-8" {
			t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
		}
	})
}

func TestSigninRoute(t *testing.T) {
	// The setupServer method, that we previously refactored
	// is injected into a test server
	ts := httptest.NewServer(server.SetupServer())
	// Shut down the server and block until all requests have gone through
	defer ts.Close()

	// Make a request to our server with the {base url}/ping
	r := strings.NewReader("any non validated body")
	resp, err := http.Post(fmt.Sprintf("%s"+server.SigninRoute, ts.URL), "application/json", r)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

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

func TestSignoutRoute(t *testing.T) {
	// The setupServer method, that we previously refactored
	// is injected into a test server
	ts := httptest.NewServer(server.SetupServer())
	// Shut down the server and block until all requests have gone through
	defer ts.Close()

	// Make a request to our server with the {base url}/ping
	r := strings.NewReader("any non validated body")
	resp, err := http.Post(fmt.Sprintf("%s"+server.SignoutRoute, ts.URL), "application/json", r)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

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
