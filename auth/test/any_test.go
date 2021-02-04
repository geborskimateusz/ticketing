package test

import (
	"testing"
)

func TestNewPersonPositiveAge(t *testing.T) {
	a := 1
	if a != 2 {
		t.Errorf("Error in test expected %v received %v", 2, a)
	}
}
