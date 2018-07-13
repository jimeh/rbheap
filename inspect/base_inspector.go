package inspect

import (
	"fmt"
	"io"
	"os"
	"time"
)

type BaseInspector struct {
	FilePath      string
	Dump          *Dump
	Verbose       bool
	VerboseWriter io.Writer
}

func (s *SourceInspector) Load() error {
	start := time.Now()
	s.verbose(fmt.Sprintf("Loading %s...", s.FilePath))

	err := s.Dump.Load()
	if err != nil {
		return err
	}

	elapsed := time.Now().Sub(start)
	s.verbose(fmt.Sprintf(
		"Loaded %d objects in %.6f seconds",
		len(s.Dump.Objects),
		elapsed.Seconds(),
	))

	return nil
}

func (s *BaseInspector) log(msg string) {
	if s.Verbose {
		fmt.Println(msg)
	}
}

func (s *BaseInspector) verbose(msg string) {
	if s.Verbose {
		w := s.VerboseWriter

		if w == nil {
			w = os.Stderr
		}

		fmt.Fprintln(w, msg)
	}
}
