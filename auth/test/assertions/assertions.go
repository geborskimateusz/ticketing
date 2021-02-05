package assertions

import "log"

func AssertStatusCode(expected, actual int) {
	if actual != expected {
		log.Fatalf("Expected status code %v, got %v", expected, actual)
	}
}
