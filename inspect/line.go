package inspect

// NewLine creates a new Line.
func NewLine(filePath string, lineNum int) *Line {
	return &Line{
		FilePath:  filePath,
		LineNum:   lineNum,
		ObjectMap: map[string]*Object{},
	}
}

// Line represents a source line within a file and the objects allocated by it.
type Line struct {
	FilePath    string
	LineNum     int
	ObjectMap   map[string]*Object
	ObjectCount int
	ByteSize    int64
	MemSize     int64
}

// Add adds a Object to a Line.
func (s *Line) Add(obj *Object) {
	_, ok := s.ObjectMap[obj.Address]
	if !ok && obj.File != "" && obj.Line != 0 {
		s.ObjectCount++
		s.ByteSize = s.ByteSize + obj.ByteSize
		s.MemSize = s.MemSize + obj.MemSize

		s.ObjectMap[obj.Address] = obj
	}
}
