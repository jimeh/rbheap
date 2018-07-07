package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func NewHeapDump(file string) (*HeapDump, error) {
	heapDump := HeapDump{File: file}
	err := heapDump.Process()
	return &heapDump, err
}

// HeapDump contains all relevant data for a single heap dump.
type HeapDump struct {
	File    string
	Index   []string
	Entries map[string]*Entry
}

// Process processes the heap dump
func (s *HeapDump) Process() error {
	file, err := os.Open(s.File)
	defer file.Close()

	if err != nil {
		return err
	}

	s.Entries = map[string]*Entry{}

	reader := bufio.NewReader(file)
	var offset int64 = -1
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
		s.Index = append(s.Index, entry.Index)
	}

	return nil
}

func (s *HeapDump) PrintMatchingJSON(indexes *[]string) error {
	file, err := os.Open(s.File)
	defer file.Close()

	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)
	offsets := s.matchingOffsets(indexes)

	var current int64 = 0
	var offset int64 = -1

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
	}

	return nil
}

func (s *HeapDump) matchingOffsets(indexes *[]string) []int64 {
	var offsets []int64

	for _, index := range *indexes {
		offsets = append(offsets, s.Entries[index].Offset)
	}

	sort.Slice(offsets, func(i, j int) bool { return offsets[i] < offsets[j] })
	return offsets
}
