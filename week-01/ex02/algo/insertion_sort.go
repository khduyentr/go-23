package algo

import (
	"github.com/khduyentr/go-23/w01/ex02/helper"
)

func InsertionSort[Type helper.AvailableType](array []Type) []Type {
	arrayLen := len(array)

	for i := 1; i < arrayLen; i++ {
		last := array[i]
		j := i

		for j > 0 && array[j-1] > last {
			array[j] = array[j-1]
			j = j - 1
		}

		array[j] = last
	}

	helper.PrintArray(array)

	return array
}

func MixInsertionSort(array []interface{}) {
	floatArray, stringArray := helper.ClassifyArrayBasedOnType(array)
	InsertionSort(floatArray)
	InsertionSort(stringArray)
}
