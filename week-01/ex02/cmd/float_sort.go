package cmd

import (
	"fmt"

	"github.com/khduyentr/go-23/w01/ex02/algo"
	"github.com/khduyentr/go-23/w01/ex02/helper"
	"github.com/spf13/cobra"
)

var floatSortCommand = &cobra.Command{
	Use:     "float",
	Aliases: []string{"-float"},
	Short:   "Sort a float64 array",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var array []float64 = helper.ConvertArrayToFloat64(args)

		fmt.Print("Output: ")
		algo.SelectionSort(array, isDescending)
		fmt.Println()

		// fmt.Print("[IS] Output: ")
		// algo.InsertionSort(array, isDescending)
		// fmt.Println()
	},
}

func init() {
	floatSortCommand.Flags().BoolVarP(&isDescending, "descending", "d", false, "Sort in descending order")
	rootCommand.AddCommand(floatSortCommand)
}
