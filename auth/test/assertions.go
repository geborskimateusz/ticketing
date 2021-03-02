package test

import (
	"io"
	"log"
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

func AssertResponseBody(t *testing.T, expected string, actual io.ReadCloser) {
	t.Helper()
	t.Fail()
}
