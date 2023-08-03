package unit_test

import (
	"testing"

	"github.com/khduyentr/go-23/w01/ex02/algo"
	"github.com/khduyentr/go-23/w01/ex02/helper"
)

func Equal[Type helper.AvailableType](a, b []Type) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestIntSelectionSort(t *testing.T) {
	data := []int64{1, 0, 3, 4, 7}

	result := algo.SelectionSort(data)

	expected := []int64{0, 1, 3, 4, 7}

	if !Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

}

func TestStringSelectionSort(t *testing.T) {
	data := []string{"apple", "dog", "cat", "bird", "egg"}

	result := algo.SelectionSort(data)

	expected := []string{"apple", "bird", "cat", "dog", "egg"}

	if !Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFloatSelectionSort(t *testing.T) {
	data := []float64{1.1, 0.5, 3.2, 4.4, 7.5}

	result := algo.SelectionSort(data)

	expected := []float64{0.5, 1.1, 3.2, 4.4, 7.5}

	if !Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestIntInsertionSort(t *testing.T) {
	data := []int64{1, 0, 3, 4, 7}

	result := algo.InsertionSort(data)

	expected := []int64{0, 1, 3, 4, 7}

	if !Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFloatInsertionSort(t *testing.T) {
	data := []float64{1.1, 0.5, 3.2, 4.4, 7.5}

	result := algo.InsertionSort(data)

	expected := []float64{0.5, 1.1, 3.2, 4.4, 7.5}

	if !Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestStringInsertionSort(t *testing.T) {
	data := []string{"apple", "dog", "cat", "bird", "egg"}

	result := algo.InsertionSort(data)

	expected := []string{"apple", "bird", "cat", "dog", "egg"}

	if !Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
