package cmd

import (
	"fmt"
	"strconv"

	"github.com/khduyentr/go-23/w01/ex02/algo"
	"github.com/spf13/cobra"
)

var intSortCommand = &cobra.Command{
	Use:     "integer",
	Aliases: []string{"int"},
	Short:   "Sort an int64 array",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// convert this args array to int64
		argsLen := len(args)
		var array []int64

		for i := 0; i < argsLen; i++ {
			str := args[i]
			n, err := strconv.ParseInt(str, 10, 64)

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
	rootCommand.AddCommand(intSortCommand)
}
