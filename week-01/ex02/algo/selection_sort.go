package algo

import (
	"fmt"
	"reflect"

	"github.com/khduyentr/go-23/w01/ex02/helper"
)

func SelectionSort[Type helper.AvailableType](array []Type, isDescending bool) {
	arrayLen := len(array)

	if !isDescending {
		for i := 0; i < arrayLen-1; i++ {
			minIndex := i
			for j := i + 1; j < arrayLen; j++ {
				if array[j] < array[minIndex] {
					minIndex = j
				}
			}

			array[i], array[minIndex] = array[minIndex], array[i]
		}
	} else {
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
	arrayLen := len(array)

	floatArray := []float64{}
	stringArray := []string{}

	for i := 0; i < arrayLen; i++ {
		switch dataType := reflect.TypeOf(array[i]).String(); dataType {
		case "int":
			floatArray = append(floatArray, float64(array[i].(int)))
		case "int64":
			floatArray = append(floatArray, float64(array[i].(int64)))
		case "float64":
			floatArray = append(floatArray, array[i].(float64))
		case "string":
			stringArray = append(stringArray, array[i].(string))
		default:
			fmt.Println("This value is not expected: ", dataType)
		}
	}

	if !isDescending {
		SelectionSort(floatArray, isDescending)
		SelectionSort(stringArray, isDescending)
	} else {
		SelectionSort(stringArray, isDescending)
		SelectionSort(floatArray, isDescending)
	}
}
