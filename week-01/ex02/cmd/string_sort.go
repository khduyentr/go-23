package cmd

import (
	"fmt"

	"github.com/khduyentr/go-23/w01/ex02/algo"
	"github.com/spf13/cobra"
)

var stringSortCommand = &cobra.Command{
	Use:     "string",
	Aliases: []string{"string"},
	Short:   "Sort a string array",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Output: ")
		algo.SelectionSort(args, isDescending)
		fmt.Println()

		// fmt.Print("[IS] Output: ")
		// algo.InsertionSort(args, isDescending)
		// fmt.Println()
	},
}

func init() {
	stringSortCommand.Flags().BoolVarP(&isDescending, "descending", "d", false, "Sort in descending order")
	rootCommand.AddCommand(stringSortCommand)
}
