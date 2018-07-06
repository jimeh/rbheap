package main

import (
	"bufio"
	json "encoding/json"
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
	File  string
	Index []string
	Items map[string]*HeapItem
}

// Process processes the heap dump
func (s *HeapDump) Process() error {
	file, err := os.Open(s.File)
	defer file.Close()

	if err != nil {
		return err
	}

	s.Items = map[string]*HeapItem{}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes(byte('\n'))
		s.ProcessLine(line)

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		return err
	}

	return nil
}

func (s *HeapDump) ProcessLine(line []byte) error {
	item, err := s.Parse(line)
	if err != nil {
		return err
	}

	if len(item.Address) > 0 {
		index := item.Address + ":" + item.Type
		s.Items[index] = item
		s.Index = append(s.Index, index)
	}
	return nil
}

func (s *HeapDump) Parse(itemJSON []byte) (*HeapItem, error) {
	var i HeapItem
	err := json.Unmarshal(itemJSON, &i)

	return &i, err
}
