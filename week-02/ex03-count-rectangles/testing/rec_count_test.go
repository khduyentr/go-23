package unit_test

import (
	"testing"

	"github.com/khduyentr/go-23/ex-03/rectangle"
)

func TestCountRectangle(t *testing.T) {
	array := [][]int{
		{1, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}

	expected := 5

	result, err := rectangle.CountRectangles(array)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("expected %d rectangles, got %d", expected, result)
	}
}
