package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// 0: file_name
	// 1: first_name
	// 2: last_name
	// from 3 -> n - 2 -> middle name
	// n - 1: country code

	var builder strings.Builder
	var middleNameBuilder strings.Builder

	argsLen := len(os.Args)

	if argsLen < 4 {
		// fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	// not used the first argument
	firstName := os.Args[1]
	lastName := os.Args[2]

	countryCode := os.Args[argsLen-1]

	for i := 3; i < argsLen-1; i++ {
		middleNameBuilder.WriteString(os.Args[i])
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
	fmt.Println(result)

}
