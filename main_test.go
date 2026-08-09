package main

import "testing"

func TestHello(t *testing.T) {
	result := hello()

	expected := "Hello, CI/CD!"

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
