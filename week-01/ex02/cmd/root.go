package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "1.0.0"

var isDescending bool
var algoName string

var rootCommand = &cobra.Command{
	Use:     "CLI Sort Algorithms",
	Version: version,
	Short:   "CLI Sort Algorithms using Golang",
	Long:    `CLI Sort Algorithms is a simple CLI for sorting things?`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
