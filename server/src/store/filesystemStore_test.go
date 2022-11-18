package store

import (
	"io"
	"os"
	"testing"

	"github.com/foxbit19/todo-app/server/src/model"
	"gotest.tools/v3/assert"
)

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
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

func TestFileSystemStore(t *testing.T) {
	t.Run("returns all todo items", func(t *testing.T) {
		database, cleanDb := createTempFile(t, `[
			{"Id": 1, "Description": "first todo", "Order": 2},
			{"Id": 2, "Description": "second todo", "Order": 3},
			{"Id": 3, "Description": "third todo", "Order": 1}
		]`)
		defer cleanDb()

		store := FileSystemStore{database}

		got := store.GetItems()
		want := []model.Item{
			{
				Id: 1,
				Description: "first todo",
				Order: 2,
			},
			{
				Id: 2,
				Description: "second todo",
				Order: 3,
			},
			{
				Id: 3,
				Description: "third todo",
				Order: 1,
			},
		}

		assert.DeepEqual(t, *got, want)
	})

	t.Run("returns a single todo item", func(t *testing.T) {
		database, cleanDb := createTempFile(t, `[
			{"Id": 1, "Description": "first todo", "Order": 2},
			{"Id": 2, "Description": "second todo", "Order": 3},
			{"Id": 3, "Description": "third todo", "Order": 1}
		]`)
		defer cleanDb()

		store := FileSystemStore{database}

		got := store.GetItem(2)

		assert.DeepEqual(t, *got, model.Item{
			Id: 2,
			Description: "second todo",
			Order: 3,
		})
	})

	/* t.Run("store a new todo items", func(t *testing.T) {
		database, buffer := []model.Item{}, new(bytes.Buffer)

		json.NewEncoder(buffer).Encode(database)

		store := FileSystemStore{buffer}

		store.StoreItem("first todo")
		got := store.GetItem(1)
		assert.DeepEqual(t, *got, database[0])
	}) */
}