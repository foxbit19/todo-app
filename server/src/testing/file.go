package common

import (
	"os"
	"testing"
)

// CreateTempFile created a temporary file inside caller's current directory
// and set a remove function to remove the file.
// It checks errors during the creation of the file and during the write action.
// The pointer of the file and the pointer to the remove function
// are returned.
func CreateTempFile(t testing.TB, initialData string) (*os.File, func()) {
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