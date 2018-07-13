package cmd

import (
	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect ObjectSpace dumps from Ruby proceees",
}

var inspectOpts = struct {
	Verbose bool
	Output  string
}{}

func init() {
	rootCmd.AddCommand(inspectCmd)

	inspectCmd.PersistentFlags().BoolVarP(&inspectOpts.Verbose,
		"verbose", "v", false,
		"print verbose information to STDERR",
	)
}
