package main

import (
	"github.com/jimeh/rbheapleak/cmd"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	cmd.Execute(&cmd.BuildInfo{
		Version: version,
		Commit:  commit,
		Date:    date,
	})
}
