package store

import "github.com/foxbit19/todo-app/server/src/model"

type InMemoryItemStore struct {
	todo []model.Item
}

func (s *InMemoryItemStore) GetItem(id int) *model.Item {
	for i := 0; i < len(s.todo); i++ {
		if(s.todo[i].Id == id) {
			return &s.todo[i]
		}
	}

	return nil
}

func (s *InMemoryItemStore) GetItems() *[]model.Item {
	return &s.todo
}

func (s *InMemoryItemStore) StoreItem(description string) {
	s.todo = append(s.todo, model.Item{
		Id: len(s.todo)+1,
		Description: description,
		Order: len(s.todo)+1,
	})
}