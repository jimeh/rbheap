package main

import (
	"encoding/json"
	"io"
	"os"
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
	Entries map[string]*HeapEntry
}

// Process processes the heap dump
func (s *HeapDump) Process() error {
	file, err := os.Open(s.File)
	defer file.Close()

	if err != nil {
		return err
	}

	s.Entries = map[string]*HeapEntry{}

	d := json.NewDecoder(file)
	for {
		var e HeapEntry
		if err := d.Decode(&e); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		index := e.Address + ":" + e.Type
		s.Entries[index] = &e
		s.Index = append(s.Index, index)
	}

	return nil
}
