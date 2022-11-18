package utils

import (
	"io"
	"testing"

	testingCommon "github.com/foxbit19/todo-app/server/src/testing"
)

func TestTape_Write(t *testing.T) {
	file, cleanFile := testingCommon.CreateTempFile(t, "my_file_data")
	defer cleanFile()

	tape := &Tape{file}

	tape.Write([]byte("123"))

	// without seek the writer remains at the same position
	// into the file. We want to read the file from the
	// beginning
	file.Seek(0, 0)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "123"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}