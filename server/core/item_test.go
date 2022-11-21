package core

import (
	"testing"

	"github.com/foxbit19/todo-app/server/model"
	"github.com/foxbit19/todo-app/server/store"
	testingCommon "github.com/foxbit19/todo-app/server/testing"
	"gotest.tools/v3/assert"
)

func TestBusinessLogic(t *testing.T) {
	var store store.ItemStore

	t.Run("it gets an item", func(t *testing.T) {
		store = testingCommon.NewStubItemStore()
		core := NewItem(store)

		got := core.Get(2)
		want := model.Item{
			Id:          2,
			Description: "this is my second todo",
			Order:       2,
		}

		assert.DeepEqual(t, *got, want)
	})

	t.Run("it gets all items", func(t *testing.T) {
		store = testingCommon.NewStubItemStore()
		core := NewItem(store)

		got := core.GetAll()
		want := []model.Item{
			{
				Id:          1,
				Description: "this is my first todo",
				Order:       1,
			},
			{
				Id:          2,
				Description: "this is my second todo",
				Order:       2,
			},
		}

		assert.DeepEqual(t, *got, want)
	})

	t.Run("it creates an item", func(t *testing.T) {
		store = testingCommon.NewStubItemStore()
		core := NewItem(store)

		got := core.Create("my new wonderful todo item")
		want := model.Item{
				Id:          3,
				Description: "my new wonderful todo item",
				Order:       3,
			}

		assert.DeepEqual(t, *got, want)
	})

	t.Run("it creates an item with a default order that is the max of existing orders", func(t *testing.T) {
		store = testingCommon.NewStubItemStore()
		store.StoreItem("my personal todo", 23)
		core := NewItem(store)

		got := core.Create("my new wonderful todo item")

		assert.Equal(t, got.Description, "my new wonderful todo item")
		assert.Equal(t, got.Order, 24)
	})

	t.Run("it updates an existing item", func(t *testing.T) {
		store = testingCommon.NewStubItemStore()
		core := NewItem(store)

		want := &model.Item{
			Id: 1,
			Description: "First todo ever!",
			Order: 1,
		}
		core.Update(want)
		got := core.Get(1)

		assert.DeepEqual(t, got, want)
	})

	t.Run("it deletes an item", func(t *testing.T) {
		store = testingCommon.NewStubItemStore()
		core := NewItem(store)

		core.Delete(2)
		got := core.Get(2)

		if got != nil {
			t.Errorf("got is not nil")
		}
	})

	t.Run("it adds priority to an item changing its order", func(t *testing.T) {
		store = testingCommon.NewStubItemStore()
		core := NewItem(store)

		core.Create("new item 1")
		core.Create("new item 2")

		source := core.Get(3)
		target := core.Get(1)

		core.Reorder(source.Id, target.Id)

		got := core.GetAll()

		want := []model.Item{
			{
				Id:          3,
				Description: "new item 1",
				Order:       1,
			},
			{
				Id:          1,
				Description: "this is my first todo",
				Order:       2,
			},
			{
				Id:          2,
				Description: "this is my second todo",
				Order:       3,
			},
			{
				Id:          4,
				Description: "new item 2",
				Order:       4,
			},
		}

		assert.DeepEqual(t, *got, want)
	})

	t.Run("it lowers the priority of an item changing its order", func(t *testing.T) {
		store = testingCommon.NewStubItemStore()
		core := NewItem(store)

		core.Create("new item 1")
		core.Create("new item 2")

		source := core.Get(2)
		target := core.Get(4)

		core.Reorder(source.Id, target.Id)

		got := core.GetAll()

		want := []model.Item{
			{
				Id:          1,
				Description: "this is my first todo",
				Order:       1,
			},
			{
				Id:          3,
				Description: "new item 1",
				Order:       2,
			},
			{
				Id:          4,
				Description: "new item 2",
				Order:       3,
			},
			{
				Id:          2,
				Description: "this is my second todo",
				Order:       4,
			},
		}

		assert.DeepEqual(t, *got, want)
	})
}