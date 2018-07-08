package leak

import (
	"fmt"
	"time"

	"github.com/jimeh/rbheap/obj"
)

func NewFinder(file1, file2, file3 string) *Finder {
	return &Finder{
		FilePaths: [3]string{file1, file2, file3},
	}
}

type Finder struct {
	FilePaths [3]string
	Dumps     [3]*obj.Dump
	Leaks     []*string
	Verbose   bool
}

func (s *Finder) Process() error {
	for i, filePath := range s.FilePaths {
		start := time.Now()
		s.log(fmt.Sprintf("Parsing %s", filePath))
		dump := obj.NewDump(filePath)

		err := dump.Process()
		if err != nil {
			return err
		}

		s.Dumps[i] = dump
		elapsed := time.Now().Sub(start)
		s.log(fmt.Sprintf(
			"Parsed %d objects in %.6f seconds",
			len(dump.Index),
			elapsed.Seconds(),
		))
	}

	return nil
}

func (s *Finder) PrintLeakedAddresses() {
	s.log("\nLeaked Addresses:")
	s.Dumps[1].PrintEntryAddress(s.FindLeaks())
}

func (s *Finder) PrintLeakedObjects() error {
	s.log("\nLeaked Objects:")
	return s.Dumps[1].PrintEntryJSON(s.FindLeaks())
}

func (s *Finder) FindLeaks() []*string {
	if s.Leaks != nil {
		return s.Leaks
	}

	mapA := map[string]bool{}
	mapC := map[string]bool{}

	for _, x := range s.Dumps[0].Index {
		mapA[*x] = true
	}

	for _, x := range s.Dumps[2].Index {
		mapC[*x] = true
	}

	for _, x := range s.Dumps[1].Index {
		_, okA := mapA[*x]
		_, okC := mapC[*x]

		if !okA && okC {
			s.Leaks = append(s.Leaks, x)
		}
	}

	return s.Leaks
}

func (s *Finder) log(msg string) {
	if s.Verbose {
		fmt.Println(msg)
	}
}
