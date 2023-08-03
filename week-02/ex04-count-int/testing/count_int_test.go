package unit_test

import (
	"testing"

	"github.com/khduyentr/go-23/ex-04/int/counter"
)

func TestCountInt(t *testing.T) {

	data := "a123bc34d8ef34"

	expected := 3

	result := counter.NumDifferentIntegers(data)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

}
