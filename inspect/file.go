package inspect

// NewFile creates a new File.
func NewFile(filePath string) *File {
	return &File{
		FilePath:  filePath,
		ObjectMap: map[string]*Object{},
	}
}

// File represents a source file and the lines and objects allocated by them.
type File struct {
	FilePath    string
	ObjectMap   map[string]*Object
	ObjectCount int
	ByteSize    int64
	MemSize     int64
}

// Add adds a object to a File.
func (s *File) Add(obj *Object) {
	_, ok := s.ObjectMap[obj.Address]
	if !ok && obj.File != "" && obj.Line != 0 {
		s.ObjectCount++
		s.ByteSize = s.ByteSize + obj.ByteSize
		s.MemSize = s.MemSize + obj.MemSize

		s.ObjectMap[obj.Address] = obj
	}
}
