package main

import (
	"fmt"
	"os"
	"strings"
)

func isUppercase(value string) bool {
	return strings.ToUpper(value) == value
}

func main() {

	var sb strings.Builder

	var first_name string = os.Args[1]
	var last_name string = os.Args[2]

	var middle_name string = ""
	var country_code string = ""
	var result string = ""

	if !isUppercase(os.Args[3]) {
		middle_name = os.Args[3]
		country_code = os.Args[4]
	} else {
		country_code = os.Args[3]
	}

	switch country_code {
	case "VN", "JP", "CN", "KR", "KH":

		sb.WriteString(last_name)

		if middle_name != "" {
			sb.WriteString(" ")
			sb.WriteString(middle_name)
		}

		sb.WriteString(" ")
		sb.WriteString(first_name)
	case "US", "UK", "AU", "FR":
	default:
		sb.WriteString(first_name)

		if middle_name != "" {
			sb.WriteString(" ")
			sb.WriteString(middle_name)
		}

		sb.WriteString(" ")
		sb.WriteString(last_name)
	}

	result = sb.String()
	fmt.Println(result)
}
