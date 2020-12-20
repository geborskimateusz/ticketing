package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/geborskimateusz/auth/api"
)

func TestSignupRoute(t *testing.T) {

	ts := httptest.NewServer(api.Instance())
	defer ts.Close()

	t.Run("should return status 200", func(t *testing.T) {

		requestBody, _ := json.Marshal(map[string]string{
			"email":    "test@gmail.com",
			"password": "123dsfsdf",
		})

		resp, err := http.Post(fmt.Sprintf("%s"+api.SignupRoute, ts.URL), "application/json", bytes.NewBuffer(requestBody))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		AssertStatusCode(t, resp.StatusCode, http.StatusOK)
		AssertHeaderAndContentType(t, resp)
	})

	t.Run("should return status on invalid email or password", func(t *testing.T) {

		requestBody, err := json.Marshal(map[string]string{
			"email":    "testgmail.com",
			"password": "123dsfsdf",
		})
		if err != nil {
			print(err)
		}
		resp, err := http.Post(fmt.Sprintf("%s"+api.SignupRoute, ts.URL), "application/json", bytes.NewBuffer(requestBody))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		AssertStatusCode(t, resp.StatusCode, http.StatusBadRequest)
		AssertHeaderAndContentType(t, resp)
		AssertValidationError(t, resp, "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag")

		requestBody, err = json.Marshal(map[string]string{
			"email":    "test@gmail.com",
			"password": "123",
		})
		if err != nil {
			print(err)
		}
		resp, err = http.Post(fmt.Sprintf("%s"+api.SignupRoute, ts.URL), "application/json", bytes.NewBuffer(requestBody))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		AssertStatusCode(t, resp.StatusCode, http.StatusBadRequest)
		AssertHeaderAndContentType(t, resp)
		AssertValidationError(t, resp, "Key: 'User.Password' Error:Field validation for 'Password' failed on the 'min' tag")

	})
}

func TestSigninRoute(t *testing.T) {
	t.Run("should return status 200", func(t *testing.T) {
		ts := httptest.NewServer(api.Instance())
		defer ts.Close()

		// Make a request to our server with the {base url}/ping
		r := strings.NewReader("any non validated body")
		resp, err := http.Post(fmt.Sprintf("%s"+api.SigninRoute, ts.URL), "application/json", r)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		AssertStatusCode(t, resp.StatusCode, http.StatusOK)
		AssertHeaderAndContentType(t, resp)
	})
}

func TestSignoutRoute(t *testing.T) {
	t.Run("should return status 200", func(t *testing.T) {
		ts := httptest.NewServer(api.Instance())
		defer ts.Close()

		// Make a request to our server with the {base url}/ping
		r := strings.NewReader("any non validated body")
		resp, err := http.Post(fmt.Sprintf("%s"+api.SignoutRoute, ts.URL), "application/json", r)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		AssertStatusCode(t, resp.StatusCode, http.StatusOK)
		AssertHeaderAndContentType(t, resp)
	})
}

func TestCurrentuserRoute(t *testing.T) {
	t.Run("should return status 200", func(t *testing.T) {
		ts := httptest.NewServer(api.Instance())
		defer ts.Close()

		// Make a request to our server with the {base url}/ping
		resp, err := http.Get(fmt.Sprintf("%s"+api.CurrentUserRoute, ts.URL))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		AssertStatusCode(t, resp.StatusCode, http.StatusOK)
		AssertHeaderAndContentType(t, resp)
	})
}
