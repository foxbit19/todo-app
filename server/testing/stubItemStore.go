package testing

import (
	"fmt"
	"sort"
	"time"

	"github.com/foxbit19/todo-app/server/helpers/slices"
	"github.com/foxbit19/todo-app/server/model"
)

type StubItemStore struct {
	Todo *[]model.Item
}

func NewStubItemStore() *StubItemStore {
	return &StubItemStore{&[]model.Item{
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
	}}
}

func (s *StubItemStore) GetItem(id int) *model.Item {
	return s.findItem(id)
}

func (s *StubItemStore) GetItems(completed bool) *[]model.Item {
	items := *slices.Filter(s.Todo, func (i int) bool  {
		return (*s.Todo)[i].Completed == completed
	})

	sort.Slice(items, func (i int, j int) bool  {
		if completed {
			// if only completed items are requestes
			// the order function works on date
			date1, _ := time.Parse(time.RFC822Z, items[i].CompletedDate)
			date2, _ := time.Parse(time.RFC822Z, items[j].CompletedDate)
			return date1.After(date2)
		} else {
			// otherwise, the order function works on order
			return items[i].Order < items[j].Order
		}
	})

	sort.Slice(items, func (i int, j int) bool  {
		return items[i].Order < items[j].Order
	})

	return &items
}

func (s *StubItemStore) StoreItem(description string, order int) int {
	id := len(*s.Todo) + 1
	todo := append(*s.Todo, model.Item{
		Id:          id,
		Description: description,
		Order:       order,
		CompletedDate: time.Now().Format(time.RFC822Z),
	})

	s.Todo = &todo

	return id
}

func (s *StubItemStore) UpdateItem(id int, item *model.Item) error {
	found := s.findItem(id)

	if found == nil {
		return fmt.Errorf("Item not found: %d", id)
	}

	found.Description = item.Description
	found.Order = item.Order
	found.Completed = item.Completed
	found.CompletedDate = item.CompletedDate

	return nil
}

func (s *StubItemStore) DeleteItem(id int) {
	index := s.findItemIndex(id)
	if index == -1 {
		return
	}
	todo := append((*s.Todo)[:index], (*s.Todo)[index+1:]...)
	s.Todo = &todo
}

func (s *StubItemStore) Reorder(itemsIds []int) {
	for index, id := range itemsIds {
		item := s.GetItem(id)
		item.Order = index + 1
		s.UpdateItem(id, item)
	}
}

func (s *StubItemStore) findItem(id int) *model.Item {
	for i := 0; i < len(*s.Todo); i++ {
		if (*s.Todo)[i].Id == id {
			return &(*s.Todo)[i]
		}
	}

	return nil
}

func (s *StubItemStore) findItemIndex(id int) int {
	for i := 0; i < len(*s.Todo); i++ {
		if (*s.Todo)[i].Id == id {
			return i
		}
	}

	return -1
}
