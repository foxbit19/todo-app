package store

import "github.com/foxbit19/todo-app/server/src/model"

type ItemStore interface {
	// GetItem gets an item from the store
	GetItem(id int) *model.Item
	// GetItems get all the items from the store
	GetItems() *[]model.Item
	// StoreItem stores an item into the store using
	// description and order
	StoreItem(description string, order int) int
	// UpdateItem updates description and order of
	// a given item. The item is found using id
	UpdateItem(id int, item *model.Item) error
	// DeleteItem deletes an item from the store using
	// its id
	DeleteItem(id int)
}