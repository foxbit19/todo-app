package utils

import (
	"os"
)

// A Tape is structure to avoid write errors
// on FileSystemStore when a write takes place.
// It contains a os.File pointer and a custom Write function
type Tape struct {
	File *os.File
}

// Write introduces a Seek(0, 0) before calling
// the write function of the file.
// So it save a file from the beginning.
func (t *Tape) Write(p []byte) (n int, err error) {
	t.File.Truncate(0)
	t.File.Seek(0, 0)
	return t.File.Write(p)
}