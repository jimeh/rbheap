package obj

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func NewDump(filePath string) *Dump {
	return &Dump{FilePath: filePath}
}

// Dump contains all relevant data for a single heap dump.
type Dump struct {
	FilePath string
	Index    []*string
	Entries  map[string]*Entry
}

// Process processes the heap dump
func (s *Dump) Process() error {
	file, err := os.Open(s.FilePath)
	defer file.Close()

	if err != nil {
		return err
	}

	s.Entries = map[string]*Entry{}

	var offset int64 = -1
	reader := bufio.NewReader(file)
	for {
		offset++
		line, err := reader.ReadBytes(byte('\n'))
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		entry, err := NewEntry(line)
		if err != nil {
			return err
		}

		entry.Offset = offset
		s.Entries[entry.Index] = entry
		s.Index = append(s.Index, &entry.Index)
	}

	return nil
}

func (s *Dump) PrintEntryAddress(indexes []*string) {
	for _, index := range indexes {
		if entry, ok := s.Entries[*index]; ok {
			fmt.Println(entry.Address())
		}
	}
}

func (s *Dump) PrintEntryJSON(indexes []*string) error {
	file, err := os.Open(s.FilePath)
	defer file.Close()

	if err != nil {
		return err
	}

	offsets := s.sortedOffsets(indexes)
	offsetsLength := int64(len(offsets))

	if offsetsLength == 0 {
		return nil
	}

	var current int64
	var offset int64 = -1
	reader := bufio.NewReader(file)
	for {
		offset++
		line, err := reader.ReadBytes(byte('\n'))
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if offset == offsets[current] {
			current++
			fmt.Print(string(line))
		}

		if current >= offsetsLength-1 {
			break
		}
	}

	return nil
}

func (s *Dump) sortedOffsets(indexes []*string) []int64 {
	var res []int64

	for _, index := range indexes {
		res = append(res, s.Entries[*index].Offset)
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })

	return res
}
