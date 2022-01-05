package main

import "testing"

func TestOla(t *testing.T) {
	result := Ola("Chris")
	expected := "Olá, Chris"

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}
