package main

import (
	"testing"
)

func TestReorderName(t *testing.T) {
	data := []string{"name", "Emily", "Rose", "Watson", "US"}

	result := ReorderNames(data)

	expected := "Emily Watson Rose"

	if result != expected {
		t.Errorf("Expected %v, got different: (%v)", expected, result)
	}
}
