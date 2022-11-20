package core

import (
	"fmt"
	"log"

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
func (i *Item) GetAll() *[]model.Item {
	return i.store.GetItems()
}

// Create creates a new item using description
func (i *Item) Create(description string) *model.Item {
	nextOrder := i.getNextOrder()
	id := i.store.StoreItem(description, nextOrder)
	return i.store.GetItem(id)
}

// Update updates the item and manage errors
func (i *Item) Update(item *model.Item) error {
	// before update the elements, we need to verify if the
	// order was changed
	currentItem := i.store.GetItem(item.Id)

	if currentItem == nil {
		return fmt.Errorf("the item to update was not found %v", item)
	}

	if item.Order >= 0 && currentItem.Order != item.Order {
		log.Printf("Change the order of the item: from %d to %d", currentItem.Order, item.Order)
		// we need to change the order
		item = i.changeOrder(currentItem, item.Order)
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

// Change the order of an item using a function to shift items
func (i *Item) changeOrder(currentItem *model.Item, order int) *model.Item {
	i.reorder(currentItem, order)

	currentItem.Order = order

	return currentItem
}

// getNextOrder is a private function to find the next order
// of the next item to insert.
// It loops through all the items looking for the max order
func (i *Item) getNextOrder() int {
	maxOrder := 0

	for _, item := range *i.store.GetItems() {
		if item.Order > maxOrder {
			maxOrder = item.Order
		}
	}

	return maxOrder+1
}

// reorder is a private function to edit the orders of all the items
// that satisfy the following conditions:
// 1 - have the order greater or equals of the new order (the items that follows the item
// 		to order when the new position takes place)
// 2 - have the order less than the current item order
func (i *Item) reorder(itemToOrder *model.Item, order int) {
	items := i.store.GetItems()

	for _, item := range *items {
		// we don't want to
		// 1-touch the order of the element to reorder
		// 2-change the order of the items that follow item to order (with the old order)
		// 3-change the order of the items with an order < than the new one
		if item.Id != itemToOrder.Id {
			if item.Order <= itemToOrder.Order && item.Order >= order {
				item.Order++
				i.store.UpdateItem(item.Id, &item)
			} else if item.Order > itemToOrder.Order && item.Order < order {
				item.Order--
				i.store.UpdateItem(item.Id, &item)
			}
		}
	}
}
