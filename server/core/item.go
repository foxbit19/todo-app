package core

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/foxbit19/todo-app/server/model"
	"github.com/foxbit19/todo-app/server/store"
)

// Item implements the business logic of
// this todo server. It contains a store to interact with
type Item struct {
	store store.ItemStore
}

// NewItem initializes the core struct using a store
func NewItem(store store.ItemStore) *Item {
	return &Item{
		store: store,
	}
}

// Get gets an item using the store
func (i *Item) Get(id int) *model.Item {
	return i.store.GetItem(id)
}

// GetAll gets all items using the store
func (i *Item) GetAll(completed bool) *[]model.Item {
	items := i.store.GetItems(completed)

	if reflect.ValueOf(*items).IsNil() {
		items = &[]model.Item{}
	}

	return items
}

// Create creates a new item using description
func (i *Item) Create(description string) *model.Item {
	nextOrder := i.getNextOrder()
	id := i.store.StoreItem(description, nextOrder)
	return i.store.GetItem(id)
}

// Update updates the item and manage errors
func (i *Item) Update(item *model.Item) error {
	currentItem := i.store.GetItem(item.Id)

	if currentItem == nil {
		return fmt.Errorf("the item to update was not found %v", item)
	}

	err := i.store.UpdateItem(item.Id, item)

	if err != nil {
		return fmt.Errorf("could not update the item %v: %q", item, err)
	}

	return nil
}

// Delete deletes an item using the func of the store
func (i *Item) Delete(id int) {
	i.store.DeleteItem(id)
}

// getNextOrder is a private function to find the next order
// of the next item to insert.
// It loops through all the items looking for the max order
func (i *Item) getNextOrder() int {
	maxOrder := 0

	for _, item := range *i.store.GetItems(false) {
		if item.Order > maxOrder {
			maxOrder = item.Order
		}
	}

	return maxOrder+1
}

func (i *Item) Reorder(sourceId int, targetId int) {
	source := i.Get(sourceId)
	target := i.Get(targetId)
	// to avoid store with memory storage
	sourceOrder := source.Order
	targetOrder := target.Order
	items := *i.GetAll(false)


	// this loop works only on the items in the reorder range:
	// from source to target elements considering order.
	// the other items are leave untouched
	for i := 0; i < len(items); i++ {
		// we don't want to touch the order of the source element
		if items[i].Id != sourceId {
			if items[i].Order < sourceOrder && items[i].Order >= target.Order {
				// Raising a priority using order:
				// we want to increase the order of an item of the list that satisfies
				// the following conditions:
				// 1 - it precedes the source item
				// 2 - it follows, or it's equals, to the target item
				items[i].Order++
			} else if items[i].Order > sourceOrder && items[i].Order <= target.Order {
				// Lowering a priority using order:
				// we want to decrease the order of an item of the list that satisfies
				// the following conditions:
				// 1 - it follows the source item
				// 2 - it precedes, or it's equals, to the target item
				items[i].Order--
			}
		} else {
			// fix source item order
			items[i].Order = targetOrder
		}
	}

	// sort all the items by order using slice function
	sort.Slice(items, func (i int, j int) bool  {
		return items[i].Order < items[j].Order
	})

	i.store.Reorder(*i.mapIds(&items))
}

func (i *Item) mapIds(items *[]model.Item) *[]int {
	ids := make([]int, len(*items))

	for i, item := range *items {
		ids[i] = item.Id
	}

	return &ids
}
