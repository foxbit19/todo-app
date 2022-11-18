package testing

import "github.com/foxbit19/todo-app/server/src/model"

type StubItemStore struct {
	Todo []model.Item
}

func (s *StubItemStore) GetItem(id int) *model.Item {
	for i := 0; i < len(s.Todo); i++ {
		if(s.Todo[i].Id == id) {
			return &s.Todo[i]
		}
	}

	return nil
}

func (s *StubItemStore) GetItems() *[]model.Item {
	return &s.Todo
}

func (s *StubItemStore) StoreItem(description string) {
	s.Todo = append(s.Todo, model.Item{
		Id: len(s.Todo)+1,
		Description: description,
		Order: 0,
	})
}
