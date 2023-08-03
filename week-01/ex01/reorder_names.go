package main

import (
	"fmt"
	"os"
	"strings"
)

func ReorderNames(data []string) string {
	var builder strings.Builder
	var middleNameBuilder strings.Builder

	argsLen := len(data)

	if argsLen < 4 {
		// fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	// not used the first argument
	firstName := data[1]
	lastName := data[2]

	countryCode := data[argsLen-1]

	for i := 3; i < argsLen-1; i++ {
		middleNameBuilder.WriteString(data[i])
		middleNameBuilder.WriteString(" ")
	}

	middleName := middleNameBuilder.String()

	switch countryCode {
	case "VN", "JP", "CN", "KR", "KH":

		builder.WriteString(lastName)

		if middleName != "" {
			builder.WriteString(" ")
			builder.WriteString(middleName)
		} else {
			builder.WriteString(" ")
		}

		builder.WriteString(firstName)
	case "US", "UK", "AU", "FR":
		builder.WriteString(firstName)

		if middleName != "" {
			builder.WriteString(" ")
			builder.WriteString(middleName)
		} else {
			builder.WriteString(" ")
		}

		builder.WriteString(lastName)
	default:
		builder.WriteString("This country code is not registered")
	}

	result := builder.String()

	return result
}

func main() {

	data := []string{"_", "Emily", "Rose", "Watson", "US"}
	// result := reorderNames(os.Args)
	result := ReorderNames(data)
	fmt.Println(data)
	fmt.Println(result)

}
