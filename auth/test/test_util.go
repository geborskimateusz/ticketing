package test

import (
	"net/http"
	"testing"
)

// AssertStatusCode check code returned from request
func AssertStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("Expected status code %v, got %v", want, got)
	}
}

// AssertHeaderAndContentType checks header and content type from request
func AssertHeaderAndContentType(t *testing.T, resp *http.Response) {
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
