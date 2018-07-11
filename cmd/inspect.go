package cmd

import (
	"fmt"

	"github.com/jimeh/rbheap/inspect"
	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect [flags] <dump-file>",
	Short: "Inspect ObjectSpace dumps from Ruby proceees",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			usage_er(cmd, fmt.Sprintf("requires 1 arg, received %d", len(args)))
		}

		inspector := inspect.New(args[0])
		inspector.Verbose = inspectOpts.Verbose
		inspector.Process()

		inspector.PrintCountByFileAndLines()
	},
}

var inspectOpts = struct {
	Verbose bool
}{}

func init() {
	rootCmd.AddCommand(inspectCmd)

	inspectCmd.PersistentFlags().BoolVarP(
		&inspectOpts.Verbose,
		"verbose", "v", false,
		"print verbose information",
	)
}
