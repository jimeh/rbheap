package inspect

import (
	"bufio"
	"io"
	"os"
)

func NewDump(filePath string) *Dump {
	return &Dump{FilePath: filePath}
}

// Dump contains all relevant data for a single heap dump.
type Dump struct {
	FilePath      string
	Objects       map[string]*Object
	ByFile        map[string][]*Object
	ByFileAndLine map[string][]*Object
	ByGeneration  map[int][]*Object
}

// Load processes the heap dump referenced in FilePath.
func (s *Dump) Load() error {
	file, err := os.Open(s.FilePath)
	defer file.Close()

	if err != nil {
		return err
	}

	s.Objects = map[string]*Object{}

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

func (s *Dump) Lookup(address string) (*Object, bool) {
	object, ok := s.Objects[address]
	return object, ok
}

// AddObject adds a *Object to the Dump.
func (s *Dump) AddObject(obj *Object) {
	s.Objects[obj.Address] = obj
}
