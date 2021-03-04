package test

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func AssertStatusCode(t *testing.T, expected, actual int) {
	t.Helper()
	if actual != expected {
		log.Fatalf("Expected status code %v, got %v", expected, actual)
	}
}

func AssertAnyError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func AssertResponseBody(t *testing.T, expected string, actual *http.Response) {
	t.Helper()

	bodyBytes, err := ioutil.ReadAll(actual.Body)
	if err != nil {
		t.Fatal(err)
	}

	actualString := string(bodyBytes)
	if actualString != expected {
		t.Fatalf("Expected response %v, got %v", expected, actualString)
	}

}
