package testing

import (
	"fmt"

	"github.com/foxbit19/todo-app/server/src/model"
)

type StubItemStore struct {
	Todo *[]model.Item
}

func (s *StubItemStore) GetItem(id int) *model.Item {
	return s.findItem(id)
}

func (s *StubItemStore) GetItems() *[]model.Item {
	return s.Todo
}

func (s *StubItemStore) StoreItem(description string) {
	todo := append(*s.Todo, model.Item{
		Id: len(*s.Todo)+1,
		Description: description,
		Order: 0,
	})

	s.Todo = &todo
}

func (s *StubItemStore) UpdateItem(id int, item *model.Item) error {
	found := s.findItem(id)

	if found == nil {
		return fmt.Errorf("Item not found: %d", id)
	}

	found.Description = item.Description
	return nil
}

func (s *StubItemStore) findItem(id int) *model.Item {
	for i := 0; i < len(*s.Todo); i++ {
		if((*s.Todo)[i].Id == id) {
			return &(*s.Todo)[i]
		}
	}

	return nil
}