package inspect

import (
	"fmt"
	"io"
	"sort"
)

// NewSourceInspector creates a new SourceInspector.
func NewSourceInspector(filePath string) *SourceInspector {
	return &SourceInspector{
		BaseInspector: BaseInspector{
			FilePath: filePath,
			Dump:     NewDump(filePath),
		},
	}
}

// SourceInspector inspects memory dumps based on the file and line that
// allocated the memory.
type SourceInspector struct {
	BaseInspector
	SortBy  string
	Limit   int
	FileMap map[string]*File
}

// ByFile writes file and line details to given io.Writer.
func (s *SourceInspector) ByFile(w io.Writer) {
	fileMap := map[string]*File{}
	files := []*File{}

	for _, obj := range s.Dump.Objects {
		if _, ok := fileMap[obj.File]; !ok {
			file := NewFile(obj.File)
			files = append(files, file)
			fileMap[obj.File] = file
		}
		fileMap[obj.File].Add(obj)
	}

	switch s.SortBy {
	case "memsize":
		sort.Slice(files, func(i, j int) bool {
			return files[i].MemSize > files[j].MemSize
		})
	case "bytesize":
		sort.Slice(files, func(i, j int) bool {
			return files[i].ByteSize > files[j].ByteSize
		})
	default:
		sort.Slice(files, func(i, j int) bool {
			return files[i].ObjectCount > files[j].ObjectCount
		})
	}

	for i, file := range files {
		fmt.Fprintf(w,
			"%s (objects: %d, bytesize: %d, memsize: %d)\n",
			file.FilePath,
			file.ObjectCount,
			file.ByteSize,
			file.MemSize,
		)

		i++
		if s.Limit != 0 && i >= s.Limit {
			break
		}
	}
}

// ByLine writes file and line details to given io.Writer.
func (s *SourceInspector) ByLine(w io.Writer) {
	lineMap := map[string]*Line{}
	lines := []*Line{}

	for _, obj := range s.Dump.Objects {
		if _, ok := lineMap[obj.File]; !ok {
			line := NewLine(obj.File, obj.Line)
			lines = append(lines, line)
			lineMap[obj.File] = line
		}
		lineMap[obj.File].Add(obj)
	}

	switch s.SortBy {
	case "memsize":
		sort.Slice(lines, func(i, j int) bool {
			return lines[i].MemSize > lines[j].MemSize
		})
	case "bytesize":
		sort.Slice(lines, func(i, j int) bool {
			return lines[i].ByteSize > lines[j].ByteSize
		})
	default:
		sort.Slice(lines, func(i, j int) bool {
			return lines[i].ObjectCount > lines[j].ObjectCount
		})
	}

	for i, file := range lines {
		fmt.Fprintf(w,
			"%s:%d (objects: %d, bytesize: %d, memsize: %d)\n",
			file.FilePath,
			file.LineNum,
			file.ObjectCount,
			file.ByteSize,
			file.MemSize,
		)

		i++
		if s.Limit != 0 && i >= s.Limit {
			break
		}
	}
}
