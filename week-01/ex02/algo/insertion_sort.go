package algo

import (
	"github.com/khduyentr/go-23/w01/ex02/helper"
)

func InsertionSort[Type helper.AvailableType](array []Type, isDescending bool) {
	arrayLen := len(array)

	if !isDescending {
		for i := 1; i < arrayLen; i++ {
			last := array[i]
			j := i

			for j > 0 && array[j-1] > last {
				array[j] = array[j-1]
				j = j - 1
			}

			array[j] = last
		}
	} else {
		for i := 1; i < arrayLen; i++ {
			last := array[i]
			j := i

			for j > 0 && array[j-1] < last {
				array[j] = array[j-1]
				j = j - 1
			}

			array[j] = last
		}
	}

	helper.PrintArray(array)
}

func MixInsertionSort(array []interface{}, isDescending bool) {
	floatArray, stringArray := helper.ClassifyArrayBasedOnType(array)
	if !isDescending {
		InsertionSort(floatArray, isDescending)
		InsertionSort(stringArray, isDescending)
	} else {
		InsertionSort(stringArray, isDescending)
		InsertionSort(floatArray, isDescending)
	}
}
