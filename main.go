package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	name    = "rbheapleak"
	version = "dev"
	commit  = "unknown"
	date    = "unknown"

	formatFlag = kingpin.Flag("format", "Output format (\"hex\" or \"full\")").
			Short('f').Default("hex").String()
	silentFlag = kingpin.Flag("silent", "Silence all info output").
			Short('s').Bool()

	file1Path = kingpin.Arg("dump-1", "Path to first heap dump file.").
			Required().String()
	file2Path = kingpin.Arg("dump-2", "Path to second heap dump file.").
			Required().String()
	file3Path = kingpin.Arg("dump-3", "Path to Third heap dump file.").
			Required().String()
)

func versionString() string {
	var buffer bytes.Buffer
	var meta []string

	buffer.WriteString(fmt.Sprintf("%s %s", name, version))

	if commit != "unknown" {
		meta = append(meta, commit)
	}

	if date != "unknown" {
		meta = append(meta, date)
	}

	if len(meta) > 0 {
		buffer.WriteString(fmt.Sprintf(" (%s)", strings.Join(meta, ", ")))
	}

	return buffer.String()
}

func main() {
	kingpin.Version(versionString())
	kingpin.Parse()

	finder := NewLeakFinder(*file1Path, *file2Path, *file3Path)
	finder.Verbose = !*silentFlag

	err := finder.Process()
	if err != nil {
		log.Fatal(err)
	}

	if *formatFlag == "hex" {
		finder.PrintLeakedAddresses()
	} else if *formatFlag == "full" {
		finder.PrintLeakedObjects()
	}
}
