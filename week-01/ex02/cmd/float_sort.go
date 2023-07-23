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
		algo.SelectionSort(array)
		fmt.Println()
	},
}

func init() {
	rootCommand.AddCommand(floatSortCommand)
}
