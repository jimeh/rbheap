package cmd

import (
	"fmt"
	"os"

	"github.com/jimeh/rbheap/inspect"
	"github.com/spf13/cobra"
)

// inspectSourceCmd represents the inspectSource command
var inspectSourceCmd = &cobra.Command{
	Use:   "source [flags] <dump-file>",
	Short: "Group objects by source filename and line number",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			usage_er(cmd, fmt.Sprintf("requires 1 arg, received %d", len(args)))
		}

		inspector := inspect.NewSourceInspector(args[0])
		inspector.Verbose = inspectOpts.Verbose
		inspector.SortBy = inspectSourceOpts.SortBy
		inspector.Limit = inspectSourceOpts.Limit
		inspector.Load()

		switch inspectSourceOpts.Breakdown {
		case "file":
			inspector.ByFile(os.Stdout)
		case "line":
			inspector.ByLine(os.Stdout)
		default:
			usage_er(cmd, "Invalid --breakdown option")
		}
	},
}

var inspectSourceOpts = struct {
	Breakdown string
	SortBy    string
	Limit     int
}{}

func init() {
	inspectCmd.AddCommand(inspectSourceCmd)

	inspectSourceCmd.Flags().StringVarP(&inspectSourceOpts.Breakdown,
		"breakdown", "b", "line",
		"Breakdown sources by \"line\" or \"file\"",
	)

	inspectSourceCmd.Flags().StringVarP(&inspectSourceOpts.SortBy,
		"sort", "s", "count",
		"Sort by \"count\", \"memsize\", or \"bytesize\"",
	)

	inspectSourceCmd.Flags().IntVarP(&inspectSourceOpts.Limit,
		"limit", "l", 0,
		"Limit number of results to show",
	)
}
