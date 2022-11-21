package store

import (
	"testing"
	"time"

	"github.com/foxbit19/todo-app/server/model"
	testingCommon "github.com/foxbit19/todo-app/server/testing"
	"gotest.tools/v3/assert"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("returns all live todo items", func(t *testing.T) {
		database, cleanDb := testingCommon.CreateTempFile(t, `[
			{"Id": 1, "Description": "first todo", "Order": 1},
			{"Id": 2, "Description": "second todo", "Order": 2},
			{"Id": 3, "Description": "third todo", "Order": 3, "Completed": true}
		]`)
		defer cleanDb()

		store, _ := NewFileSystemStore(database)

		got := store.GetItems(false)
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
		}

		assert.DeepEqual(t, *got, want)
	})

	t.Run("returns all completed todo items", func(t *testing.T) {
		database, cleanDb := testingCommon.CreateTempFile(t, `[
			{"Id": 1, "Description": "first todo", "Order": 1},
			{"Id": 2, "Description": "second todo", "Order": 2},
			{"Id": 3, "Description": "third todo", "Order": 3, "Completed": true}
		]`)
		defer cleanDb()

		store, _ := NewFileSystemStore(database)

		got := store.GetItems(true)
		want := []model.Item{
			{
				Id: 3,
				Description: "third todo",
				Order: 3,
				Completed: true,
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

		store.StoreItem("first todo", 1)
		got := store.GetItem(1)
		want := model.Item{
			Id: 1,
			Description: "first todo",
			Order: 1,
			Completed: false,
		}
		assert.Equal(t, got.Id, want.Id)
		assert.Equal(t, got.Description, want.Description)
		assert.Equal(t, got.Order, want.Order)
		assert.Equal(t, got.Completed, want.Completed)
		// completed date default value needs to be in the correct format (RFC822Z)
		_, err := time.Parse(time.RFC822Z,got.CompletedDate)
		assert.NilError(t, err)

	})

	t.Run("Gives an error when database file is empty", func(t *testing.T) {
		database, cleanDb := testingCommon.CreateTempFile(t, "")
		defer cleanDb()

		_, err := NewFileSystemStore(database)

		assert.NilError(t, err)
	})

	t.Run("returns the todo list in order when live items are returned", func(t *testing.T) {
		database, cleanDb := testingCommon.CreateTempFile(t, `[
			{"Id": 1, "Description": "first todo", "Order": 2},
			{"Id": 2, "Description": "second todo", "Order": 3},
			{"Id": 3, "Description": "third todo", "Order": 1}
		]`)
		defer cleanDb()

		store, _ := NewFileSystemStore(database)

		got := store.GetItems(false)
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

	t.Run("returns the todo list using completed date in descending order when completed items are returned", func(t *testing.T) {
		database, cleanDb := testingCommon.CreateTempFile(t, `[
			{"Id": 1, "Description": "first todo", "Order": 2, "Completed": true, "CompletedDate": "20 Nov 22 15:04 +0100"},
			{"Id": 2, "Description": "second todo", "Order": 3, "Completed": true, "CompletedDate": "21 Nov 22 11:00 +0100"},
			{"Id": 3, "Description": "third todo", "Order": 1, "Completed": true, "CompletedDate": "19 Nov 22 14:30 +0100"}
		]`)
		defer cleanDb()

		store, _ := NewFileSystemStore(database)

		got := store.GetItems(true)

		want := []model.Item{
			{
				Id: 2,
				Description: "second todo",
				Order: 3,
				Completed: true,
				CompletedDate: "21 Nov 22 11:00 +0100",
			},
			{
				Id: 1,
				Description: "first todo",
				Order: 2,
				Completed: true,
				CompletedDate: "20 Nov 22 15:04 +0100",
			},
			{
				Id: 3,
				Description: "third todo",
				Order: 1,
				Completed: true,
				CompletedDate: "19 Nov 22 14:30 +0100",
			},
		}

		assert.DeepEqual(t, *got, want)
	})

	t.Run("it updates the description of an existing todo item", func(t *testing.T) {
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

	t.Run("it updates the order of an existing todo item", func(t *testing.T) {
		database, cleanDb := testingCommon.CreateTempFile(t, `[
			{"Id": 3, "Description": "third todo", "Order": 1}
		]`)
		defer cleanDb()

		store, _ := NewFileSystemStore(database)

		want := 5
		store.UpdateItem(3, &model.Item{
			Description: "third todo",
			Order: want,
		})
		got := store.GetItem(3)

		assert.Equal(t, got.Order, want)
	})

	t.Run("it updates completed info about an existing todo item", func(t *testing.T) {
		database, cleanDb := testingCommon.CreateTempFile(t, `[
			{"Id": 3, "Description": "third todo", "Order": 1, "Completed": false, "CompletedDate": "19 Nov 22 14:30 +0100"}
		]`)
		defer cleanDb()

		store, _ := NewFileSystemStore(database)

		store.UpdateItem(3, &model.Item{
			Description: "third todo",
			Order: 1,
			Completed: true,
			CompletedDate: "21 Nov 22 14:30 +0100",
		})
		got := store.GetItem(3)

		assert.Equal(t, got.Completed, true)
		assert.Equal(t, got.CompletedDate, "21 Nov 22 14:30 +0100")
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

	t.Run("it deletes an existing item", func(t *testing.T) {
		database, cleanDb := testingCommon.CreateTempFile(t, `[
			{"Id": 3, "Description": "third todo", "Order": 1}
		]`)
		defer cleanDb()

		store, _ := NewFileSystemStore(database)

		store.DeleteItem(3)
		got := store.GetItem(3)

		if got != nil {
			t.Errorf("got %v is not nil", got)
		}
	})

	t.Run("it reorder the items", func(t *testing.T) {
		database, cleanDb := testingCommon.CreateTempFile(t, `[]`)
		defer cleanDb()

		store, _ := NewFileSystemStore(database)

		store.StoreItem("first todo", 1)
		store.StoreItem("second todo", 1)

		store.Reorder([]int{2,1})
		got := *store.GetItems(false)

		assert.Equal(t, got[0].Id, 2)
		assert.Equal(t, got[1].Id, 1)
	})


}