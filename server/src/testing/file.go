package common

import (
	"io"
	"os"
	"testing"
)

// CreateTempFile created a temporary file inside caller's current directory
// and set a remove function to remove the file.
// It checks errors during the creation of the file and during the write action.
// The handler of the file using a ReadWriteSeeker and the pointer to the function
// to remove the file are returned.
func CreateTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("./", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	_, error := tmpfile.Write([]byte(initialData))

	if error != nil {
		t.Fatalf("could not write on temp file %v", error)
	}

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}