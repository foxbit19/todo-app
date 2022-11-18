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
			{"Id": 1, "Description": "first todo", "Order": 1},
			{"Id": 2, "Description": "second todo", "Order": 2},
			{"Id": 3, "Description": "third todo", "Order": 3}
		]`)
		defer cleanDb()

		store, _ := NewFileSystemStore(database)

		got := store.GetItems()
		want := []model.Item{
			{
				Id: 1,
				Description: "first todo",
				Order: 1,
			},
			{
				Id: 2,
				Description: "second todo",
				Order: 2,
			},
			{
				Id: 3,
				Description: "third todo",
				Order: 3,
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

		store, _ := NewFileSystemStore(database)

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

		store, _ := NewFileSystemStore(database)

		store.StoreItem("first todo")
		got := store.GetItem(1)
		assert.DeepEqual(t, *got, model.Item{
			Id: 1,
			Description: "first todo",
			Order: 1,
		})
	})

	t.Run("Gives an error when database file is empty", func(t *testing.T) {
		database, cleanDb := testingCommon.CreateTempFile(t, "")
		defer cleanDb()

		_, err := NewFileSystemStore(database)

		assert.NilError(t, err)
	})

	t.Run("returns the todo list in order", func(t *testing.T) {
		database, cleanDb := testingCommon.CreateTempFile(t, `[
			{"Id": 1, "Description": "first todo", "Order": 2},
			{"Id": 2, "Description": "second todo", "Order": 3},
			{"Id": 3, "Description": "third todo", "Order": 1}
		]`)
		defer cleanDb()

		store, _ := NewFileSystemStore(database)

		got := store.GetItems()
		want := []model.Item{
			{
				Id: 3,
				Description: "third todo",
				Order: 1,
			},
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
		}

		assert.DeepEqual(t, *got, want)
	})

	t.Run("updated an existing todo item", func(t *testing.T) {
		database, cleanDb := testingCommon.CreateTempFile(t, `[
			{"Id": 3, "Description": "third todo", "Order": 1}
		]`)
		defer cleanDb()

		store, _ := NewFileSystemStore(database)

		want := "new description here"
		store.UpdateItem(3, &model.Item{
			Description: want,
		})
		got := store.GetItem(3)

		assert.DeepEqual(t, got.Description, want)
	})

	t.Run("it gives an error when trying to updated an non-existing todo item", func(t *testing.T) {
		database, cleanDb := testingCommon.CreateTempFile(t, `[
			{"Id": 3, "Description": "third todo", "Order": 1}
		]`)
		defer cleanDb()

		store, _ := NewFileSystemStore(database)

		want := "new description here"
		err := store.UpdateItem(5, &model.Item{
			Description: want,
		})

		assert.Error(t, err, "Item 5 not found")
	})

	t.Run("it gives an error when trying to updated an item without description", func(t *testing.T) {
		database, cleanDb := testingCommon.CreateTempFile(t, `[
			{"Id": 3, "Description": "third todo", "Order": 1}
		]`)
		defer cleanDb()

		store, _ := NewFileSystemStore(database)

		err := store.UpdateItem(3, &model.Item{
			Description: "",
		})

		assert.Error(t, err, "Cannot update item 3 without a description")
	})
}