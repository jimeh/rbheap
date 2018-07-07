package main

import (
	"bytes"
	"fmt"
	"log"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	name    = "rbheapleak"
	version = "dev"
	commit  = "unknown"
	date    = "unknown"

	formatFlag = kingpin.Flag("format", "Output format (\"hex\" or \"full\")").
			Default("hex").String()

	file1Path = kingpin.Arg("dump-1", "Path to first heap dump file.").
			Required().String()
	file2Path = kingpin.Arg("dump-2", "Path to second heap dump file.").
			Required().String()
	file3Path = kingpin.Arg("dump-3", "Path to Third heap dump file.").
			Required().String()
)

func versionString() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s %s", name, version))

	if commit != "unknown" {
		buffer.WriteString(fmt.Sprintf(" (%s", commit))
		if date != "unknown" {
			buffer.WriteString(fmt.Sprintf(" %s", date))
		}
		buffer.WriteString(")")
	}

	return buffer.String()
}

// diffsect removes all items in `a` from `b`, then removes all items from `b`
// which are not in `c`. Effectively: intersect(difference(b, a), c)
func diffsect(a, b, c *[]string) *[]string {
	result := []string{}
	mapA := map[string]bool{}
	mapC := map[string]bool{}

	for _, x := range *a {
		mapA[x] = true
	}

	for _, x := range *c {
		mapC[x] = true
	}

	for _, x := range *b {
		_, okA := mapA[x]
		_, okC := mapC[x]

		if !okA && okC {
			result = append(result, x)
		}
	}

	return &result
}

func printHexDiff(leaked *[]string, dump *HeapDump) {
	for _, index := range *leaked {
		if item, ok := dump.Entries[index]; ok {
			fmt.Printf("%s\n", item.Address)
		}
	}
}

func main() {
	kingpin.Version(versionString())
	kingpin.Parse()

	fmt.Printf("--> Loading %s...\n", *file1Path)
	dump1, err := NewHeapDump(*file1Path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("    Loaded %d addresses\n", len(dump1.Index))

	fmt.Printf("--> Loading %s...\n", *file2Path)
	dump2, err := NewHeapDump(*file2Path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("    Loaded %d addresses\n", len(dump2.Index))

	fmt.Printf("--> Loading %s...\n", *file3Path)
	dump3, err := NewHeapDump(*file3Path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("    Loaded %d addresses\n", len(dump3.Index))

	leaked := diffsect(&dump1.Index, &dump2.Index, &dump3.Index)

	if *formatFlag == "hex" {
		printHexDiff(leaked, dump2)
	} //  else if *formatFlag == 'full' {
	//		printFullDiff(leaked, dump2)
	// }
}
