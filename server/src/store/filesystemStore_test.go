package store

import (
	"testing"

	"github.com/foxbit19/todo-app/server/src/model"
	testingCommon "github.com/foxbit19/todo-app/server/src/testing"
	"gotest.tools/v3/assert"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("returns all todo items", func(t *testing.T) {
		database, cleanDb := testingCommon.CreateTempFile(t, `[
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
		database, cleanDb := testingCommon.CreateTempFile(t, `[
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

	t.Run("store a new todo items", func(t *testing.T) {
		database, cleanDb := testingCommon.CreateTempFile(t, `[]`)
		defer cleanDb()

		store := FileSystemStore{database}

		store.StoreItem("first todo")
		got := store.GetItem(1)
		assert.DeepEqual(t, *got, model.Item{
			Id: 1,
			Description: "first todo",
			Order: 1,
		})
	})
}