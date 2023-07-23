package cmd

import (
	"fmt"

	"github.com/khduyentr/go-23/w01/ex02/algo"
	"github.com/khduyentr/go-23/w01/ex02/helper"
	"github.com/spf13/cobra"
)

var intSortCommand = &cobra.Command{
	Use:     "integer",
	Aliases: []string{"int"},
	Short:   "Sort an int64 array",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var array []int64 = helper.ConvertArrayToInt64(args)

		fmt.Print("Output: ")
		algo.SelectionSort(array, isDescending)
		fmt.Println()

		// fmt.Print("[IS] Output: ")
		// algo.InsertionSort(array, isDescending)
		// fmt.Println()
	},
}

func init() {
	intSortCommand.Flags().BoolVarP(&isDescending, "descending", "d", false, "Sort in descending order")
	rootCommand.AddCommand(intSortCommand)
}
