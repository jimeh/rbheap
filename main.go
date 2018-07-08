package main

import (
	"github.com/jimeh/rbheap/cmd"
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
