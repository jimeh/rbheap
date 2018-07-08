package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func NewObjectDump(file string) *ObjectDump {
	return &ObjectDump{File: file}
}

// ObjectDump contains all relevant data for a single heap dump.
type ObjectDump struct {
	File    string
	Index   []*string
	Entries map[string]*Entry
}

// Process processes the heap dump
func (s *ObjectDump) Process() error {
	file, err := os.Open(s.File)
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

func (s *ObjectDump) PrintEntryAddress(indexes []*string) {
	for _, index := range indexes {
		if entry, ok := s.Entries[*index]; ok {
			fmt.Println(entry.Address())
		}
	}
}

func (s *ObjectDump) PrintEntryJSON(indexes []*string) error {
	file, err := os.Open(s.File)
	defer file.Close()

	if err != nil {
		return err
	}

	offsets := s.sortedOffsets(indexes)
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
	}

	return nil
}

func (s *ObjectDump) sortedOffsets(indexes []*string) []int64 {
	var res []int64

	for _, index := range indexes {
		res = append(res, s.Entries[*index].Offset)
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })

	return res
}
