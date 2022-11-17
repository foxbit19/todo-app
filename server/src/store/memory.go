package store

import "github.com/foxbit19/todo-app/server/src/model"

type InMemoryItemStore struct {}

func (s *InMemoryItemStore) GetItem(id int) *model.Item {
	panic("No implementation here")
}

func (s *InMemoryItemStore) StoreItem(description string) {
	panic("No implementation here")
}