package cmd

import (
	"fmt"
	"strconv"

	"github.com/khduyentr/go-23/w01/ex02/algo"
	"github.com/spf13/cobra"
)

var mixSortCommand = &cobra.Command{
	Use:     "mix",
	Aliases: []string{"mix"},
	Short:   "Sort an array which has various datatypes",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		argsLen := len(args)
		var array []interface{}

		for i := 0; i < argsLen; i++ {
			temp := args[i]

			if isInt64(temp) {
				intElement, _ := strconv.ParseInt(temp, 10, 64)
				array = append(array, intElement)
			} else if isFloat64(temp) {
				floatElement, _ := strconv.ParseFloat(temp, 64)
				array = append(array, floatElement)
			} else {
				array = append(array, temp)
			}
		}

		fmt.Print("Output: ")
		algo.MixSelectionSort(array)
		fmt.Println()
	},
}

func isInt64(data string) bool {
	_, err := strconv.ParseInt(data, 10, 64)

	if err == nil {
		return true
	} else {
		return false
	}
}

func isFloat64(data string) bool {
	_, err := strconv.ParseFloat(data, 64)

	if err == nil {
		return true
	} else {
		return false
	}
}

func init() {
	rootCommand.AddCommand(mixSortCommand)
}
