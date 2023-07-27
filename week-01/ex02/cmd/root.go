package cmd

import (
	"flag"
	"fmt"

	"github.com/khduyentr/go-23/w01/ex02/algo"
	"github.com/khduyentr/go-23/w01/ex02/helper"
)

var (
	intData    bool
	floatData  bool
	stringData bool
	mixData    bool
)

func init() {

	flag.BoolVar(&intData, "int", false, "Array is all integer elements")
	flag.BoolVar(&intData, "i", false, "Array is all integer - shorthand cmd")

	flag.BoolVar(&floatData, "float", false, "Array is all float elements")
	flag.BoolVar(&floatData, "f", false, "Array is all float elements - shorthand cmd")

	flag.BoolVar(&stringData, "string", false, "Array is all string elements")
	flag.BoolVar(&stringData, "s", false, "Array is all string elements - shorthand cmd")

	flag.BoolVar(&mixData, string("mix"), false, "Array has mix datatype")
}

func Execute() {

	flag.Parse()
	args := flag.Args()

	if intData {
		var array []int64 = helper.ConvertArrayToInt64(args)

		fmt.Print("Output: ")
		algo.SelectionSort(array)
		fmt.Println()
	} else if floatData {
		var array []float64 = helper.ConvertArrayToFloat64(args)

		fmt.Print("Output: ")
		algo.SelectionSort(array)
		fmt.Println()
	} else if stringData {
		fmt.Print("Output: ")
		algo.SelectionSort(args)
		fmt.Println()
	} else if mixData {
		var array []interface{} = helper.ConvertArrayToExactType(args)

		fmt.Print("Output: ")
		algo.MixSelectionSort(array)
		fmt.Println()
	}
}
