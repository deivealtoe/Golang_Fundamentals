package main

import "testing"

func TestOla(t *testing.T) {
	result := Ola()
	expected := "Hello, World!"

	if result != expected {
		t.Errorf("Result '%s', expected '%s'", result, expected)
	}
}
