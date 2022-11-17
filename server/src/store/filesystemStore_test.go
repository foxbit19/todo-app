package store

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/foxbit19/todo-app/server/src/model"
	"gotest.tools/v3/assert"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("returns all todo items", func(t *testing.T) {
		database, buffer := []model.Item{
			{Id: 1, Description: "first todo", Order: 2},
			{Id: 2, Description: "second todo", Order: 3},
			{Id: 3, Description: "third todo", Order: 1},
		}, new(bytes.Buffer)

		json.NewEncoder(buffer).Encode(database)

		store := FileSystemStore{buffer}

		got := store.GetItems()

		assert.DeepEqual(t, *got, database)
	})
}