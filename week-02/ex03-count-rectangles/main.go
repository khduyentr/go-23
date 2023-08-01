package main

import (
	"fmt"

	"github.com/khduyentr/go-23/ex-03/rectangle"
)

func main() {

	array := [][]int{
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}

	count, err := rectangle.CountRectangles(array)

	if err != nil {
		fmt.Printf("")
	}

	fmt.Println(count)
}
