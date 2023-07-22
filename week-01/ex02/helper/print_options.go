package helper

import (
	"fmt"
)

func PrintResult[Type AvailableType](array []Type) {
	arrayLen := len(array)

	for i := 0; i < arrayLen; i++ {
		fmt.Print(array[i])
		fmt.Print(" ")
	}
}
