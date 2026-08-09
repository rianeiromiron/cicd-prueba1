package main

import "testing"

func TestHello(t *testing.T) {
	result := hello()

	expected := "Hello, world!"

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
