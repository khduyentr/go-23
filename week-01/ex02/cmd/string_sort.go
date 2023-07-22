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
		algo.SelectionSort(args)
		fmt.Println()
	},
}

func init() {
	rootCommand.AddCommand(stringSortCommand)
}
