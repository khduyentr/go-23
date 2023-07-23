package helper

import (
	"fmt"
	"reflect"
	"strconv"
)

func ConvertArrayToInt64(array []string) (result []int64) {
	arrayLength := len(array)

	for i := 0; i < arrayLength; i++ {
		str := array[i]
		n, err := strconv.ParseInt(str, 10, 64)

		if err != nil {
			panic(err)
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

func ClassifyArrayBasedOnType(array []interface{}) (floatArray []float64, stringArray []string) {

	arrayLen := len(array)

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

			return nil, nil
		}
	}

	return floatArray, stringArray
}
