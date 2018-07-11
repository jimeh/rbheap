package inspect

import (
	"fmt"
	"time"
)

func New(filePath string) *Inspector {

	return &Inspector{
		FilePath: filePath,
		Dump:     NewDump(filePath),
	}
}

type Inspector struct {
	FilePath string
	Dump     *Dump
	Verbose  bool
}

func (s *Inspector) Process() {

	start := time.Now()
	s.log(fmt.Sprintf("Parsing %s", s.FilePath))

	s.Dump.Process()

	elapsed := time.Now().Sub(start)
	s.log(fmt.Sprintf(
		"Parsed %d objects in %.6f seconds",
		len(s.Dump.ByAddress),
		elapsed.Seconds(),
	))
}

func (s *Inspector) PrintCountByFileAndLines() {
	for k, objects := range s.Dump.ByFileAndLine {
		fmt.Printf("%s: %d objects\n", k, len(objects))
	}
}

func (s *Inspector) log(msg string) {
	if s.Verbose {
		fmt.Println(msg)
	}
}
