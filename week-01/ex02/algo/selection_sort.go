package algo

import (
	"github.com/khduyentr/go-23/w01/ex02/helper"
)

func SelectionSort[Type helper.AvailableType](array []Type) []Type {
	arrayLen := len(array)

	for i := 0; i < arrayLen-1; i++ {
		minIndex := i
		for j := i + 1; j < arrayLen; j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}

		array[i], array[minIndex] = array[minIndex], array[i]
	}

	helper.PrintArray(array)

	return array
}

func MixSelectionSort(array []interface{}) {

	floatArray, stringArray := helper.ClassifyArrayBasedOnType(array)

	SelectionSort(floatArray)
	SelectionSort(stringArray)
}
