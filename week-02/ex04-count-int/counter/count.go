package counter

import (
	"strconv"
	"strings"
)

func NumDifferentIntegers(word string) int {
	var numberArrays []string
	array := strings.Split(word, "")

	for index := 0; index < len(array); index++ {
		if _, err := strconv.ParseInt(array[index], 10, 64); err == nil {
			number, currentIndex := findNumberFromIndex(array, index)

			if len(numberArrays) == 0 {
				numberArrays = append(numberArrays, number)
			} else {
				if duplicated := checkDuplication(array, number); !duplicated {
					numberArrays = append(numberArrays, number)
				}
			}

			index = currentIndex
		}
	}
	return len(numberArrays)
}

func checkDuplication(array []string, value string) bool {

	for _, val := range array {
		if val == value {
			return true
		}
	}

	return false
}

func findNumberFromIndex(array []string, index int) (string, int) {

	number := string(array[index])

	for i := index + 1; i < len(array); i++ {

		digit, err := strconv.ParseInt(array[i], 10, 64)
		if err == nil {
			number += strconv.FormatInt(int64(digit), 10)
		} else {
			return omitLeftSideZeros(number), i // result and current index
		}
	}
	return omitLeftSideZeros(number), len(array) - 1 // result and current index
}

func omitLeftSideZeros(value string) string {
	var output string
	leadingZero := true
	for _, ch := range value {
		if ch != '0' || !leadingZero {
			output += string(ch)
			leadingZero = false
		}
	}
	return output
}
