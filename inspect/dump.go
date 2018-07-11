package inspect

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func NewDump(filePath string) *Dump {
	return &Dump{FilePath: filePath}
}

// Dump contains all relevant data for a single heap dump.
type Dump struct {
	FilePath      string
	ByAddress     map[string]*Object
	ByFile        map[string][]*Object
	ByFileAndLine map[string][]*Object
	ByGeneration  map[int][]*Object
}

// Process processes the heap dump referenced in FilePath.
func (s *Dump) Process() error {
	file, err := os.Open(s.FilePath)
	defer file.Close()

	if err != nil {
		return err
	}

	s.ByAddress = map[string]*Object{}
	s.ByFile = map[string][]*Object{}
	s.ByFileAndLine = map[string][]*Object{}
	s.ByGeneration = map[int][]*Object{}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes(byte('\n'))
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		object, err := NewObject(line)
		if err != nil {
			return err
		}

		s.AddObject(object)
	}

	return nil
}

// AddObject adds a *Object to the Dump.
func (s *Dump) AddObject(obj *Object) {
	s.ByAddress[obj.Address] = obj

	if obj.File != "" {
		s.ByFile[obj.File] = append(s.ByFile[obj.File], obj)
	}

	if obj.File != "" && obj.Line != 0 {
		key := obj.File + ":" + strconv.Itoa(obj.Line)
		s.ByFileAndLine[key] = append(s.ByFileAndLine[key], obj)
	}

	if obj.Generation != 0 {
		s.ByGeneration[obj.Generation] =
			append(s.ByGeneration[obj.Generation], obj)
	}
}
