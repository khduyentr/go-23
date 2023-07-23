package algo

import (
	"github.com/khduyentr/go-23/w01/ex02/helper"
)

func SelectionSort[Type helper.AvailableType](array []Type, isDescending bool) {
	arrayLen := len(array)

	if !isDescending { // ascending
		for i := 0; i < arrayLen-1; i++ {
			minIndex := i
			for j := i + 1; j < arrayLen; j++ {
				if array[j] < array[minIndex] {
					minIndex = j
				}
			}

			array[i], array[minIndex] = array[minIndex], array[i]
		}
	} else { // descending
		for i := 0; i < arrayLen-1; i++ {
			maxIndex := i
			for j := i + 1; j < arrayLen; j++ {
				if array[j] > array[maxIndex] {
					maxIndex = j
				}
			}

			array[i], array[maxIndex] = array[maxIndex], array[i]
		}
	}

	helper.PrintArray(array)
}

func MixSelectionSort(array []interface{}, isDescending bool) {

	floatArray, stringArray := helper.ClassifyArrayBasedOnType(array)

	if !isDescending {
		SelectionSort(floatArray, isDescending)
		SelectionSort(stringArray, isDescending)
	} else {
		SelectionSort(stringArray, isDescending)
		SelectionSort(floatArray, isDescending)
	}
}
