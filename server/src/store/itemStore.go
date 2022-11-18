package store

import "github.com/foxbit19/todo-app/server/src/model"

type ItemStore interface {
	GetItem(id int) *model.Item
	GetItems() *[]model.Item
	StoreItem(description string)
	UpdateItem(id int, item *model.Item) error
	DeleteItem(id int)
}