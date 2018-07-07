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

func printHexDiff(leaked *[]string, dump *HeapDump) {
	for _, index := range *leaked {
		if item, ok := dump.Entries[index]; ok {
			fmt.Printf("%s\n", item.Address)
		}
	}
}

func logMsg(msg string) {
	if !*silentFlag {
		fmt.Println(msg)
	}
}

func loadDump(filePath string) (*HeapDump, error) {
	logMsg(fmt.Sprintf("--> Loading %s...", filePath))
	dump, err := NewHeapDump(filePath)
	logMsg(fmt.Sprintf("    Loaded %d addresses", len(dump.Index)))
	return dump, err
}

func main() {
	kingpin.Version(versionString())
	kingpin.Parse()

	dump1, err := loadDump(*file1Path)
	if err != nil {
		log.Fatal(err)
	}

	dump2, err := loadDump(*file2Path)
	if err != nil {
		log.Fatal(err)
	}

	dump3, err := loadDump(*file3Path)
	if err != nil {
		log.Fatal(err)
	}

	leaked := diffsect(&dump1.Index, &dump2.Index, &dump3.Index)

	if *formatFlag == "hex" {
		printHexDiff(leaked, dump2)
	} //  else if *formatFlag == 'full' {
	//		printFullDiff(leaked, dump2)
	// }
}
