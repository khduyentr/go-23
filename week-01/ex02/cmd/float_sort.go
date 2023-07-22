package cmd

import (
	"fmt"
	"strconv"

	"github.com/khduyentr/go-23/w01/ex02/algo"
	"github.com/spf13/cobra"
)

var floatSortCommand = &cobra.Command{
	Use:     "float",
	Aliases: []string{"-float"},
	Short:   "Sort a float64 array",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		argsLen := len(args)
		var array []float64

		for i := 0; i < argsLen; i++ {
			str := args[i]
			n, err := strconv.ParseFloat(str, 64)

			if err != nil {
				panic(err)
			}

			array = append(array, n)
		}

		fmt.Print("Output: ")
		algo.SelectionSort(array)
		fmt.Println()
	},
}

func init() {
	rootCommand.AddCommand(floatSortCommand)
}
