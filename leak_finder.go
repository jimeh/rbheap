package main

import "fmt"

func NewLeakFinder(file1, file2, file3 string) *LeakFinder {
	return &LeakFinder{
		FilePaths: [3]string{file1, file2, file3},
	}
}

type LeakFinder struct {
	FilePaths [3]string
	Dumps     [3]*Dump
	Leaks     []*string
	Verbose   bool
}

func (s *LeakFinder) Process() error {
	for i, filePath := range s.FilePaths {
		s.log(fmt.Sprintf("Parsing %s", filePath))
		dump := NewDump(filePath)

		err := dump.Process()
		if err != nil {
			return err
		}

		s.Dumps[i] = dump
		s.log(fmt.Sprintf("Parsed %d objects", len(dump.Index)))
	}

	return nil
}

func (s *LeakFinder) PrintLeakedAddresses() {
	s.log("\nLeaked Addresses:")
	s.Dumps[1].PrintEntryAddress(s.FindLeaks())
}

func (s *LeakFinder) PrintLeakedObjects() {
	s.log("\nLeaked Objects:")
	s.Dumps[1].PrintEntryJSON(s.FindLeaks())
}

func (s *LeakFinder) FindLeaks() []*string {
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

func (s *LeakFinder) log(msg string) {
	if s.Verbose {
		fmt.Println(msg)
	}
}
