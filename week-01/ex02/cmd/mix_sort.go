package cmd

import (
	"fmt"

	"github.com/khduyentr/go-23/w01/ex02/algo"
	"github.com/khduyentr/go-23/w01/ex02/helper"
	"github.com/spf13/cobra"
)

var mixSortCommand = &cobra.Command{
	Use:     "mix",
	Aliases: []string{"mix"},
	Short:   "Sort an array which has various datatypes",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var array []interface{} = helper.ConvertArrayToExactType(args)

		fmt.Print("Output: ")
		algo.MixSelectionSort(array, isDescending)
		fmt.Println()
	},
}

func init() {
	mixSortCommand.Flags().BoolVarP(&isDescending, "descending", "d", false, "Sort in descending order")
	rootCommand.AddCommand(mixSortCommand)
}
