package helper

import (
	"strconv"
)

func ConvertArrayToInt64(array []string) (result []int64) {
	arrayLength := len(array)

	for i := 0; i < arrayLength; i++ {
		str := array[i]
		n, err := strconv.ParseInt(str, 10, 64)

		if err != nil {
			panic(err)
			//return nil
		}

		result = append(result, n)
	}

	return result
}

func ConvertArrayToFloat64(array []string) (result []float64) {
	arrayLength := len(array)

	for i := 0; i < arrayLength; i++ {
		str := array[i]
		n, err := strconv.ParseFloat(str, 64)

		if err != nil {
			panic(err)
			// return nil
		}

		result = append(result, n)
	}

	return result
}

func ConvertArrayToExactType(array []string) (result []interface{}) {
	arrayLen := len(array)
	for i := 0; i < arrayLen; i++ {
		temp := array[i]

		if isInt64(temp) {
			intElement, _ := strconv.ParseInt(temp, 10, 64)
			result = append(result, intElement)
		} else if isFloat64(temp) {
			floatElement, _ := strconv.ParseFloat(temp, 64)
			result = append(result, floatElement)
		} else {
			result = append(result, temp)
		}
	}

	return result
}
