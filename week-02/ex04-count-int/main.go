package main

import (
	"fmt"

	"github.com/khduyentr/go-23/ex-04/int/counter"
)

func main() {
	word := "a123bc34d8ef34"
	// word := "A1b01c001"
	count := counter.NumDifferentIntegers(word)
	fmt.Println(count)
}
