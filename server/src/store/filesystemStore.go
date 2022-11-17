package store

import (
	"encoding/json"
	"io"

	"github.com/foxbit19/todo-app/server/src/model"
)

type FileSystemStore struct {
	database io.Reader
}

func (s *FileSystemStore) GetItem(id int) *model.Item {
	/* for i := 0; i < len(s.todo); i++ {
		if(s.todo[i].Id == id) {
			return &s.todo[i]
		}
	}

	return nil */
	return nil
}

func (s *FileSystemStore) GetItems() *[]model.Item {
	var items []model.Item
	json.NewDecoder(s.database).Decode(&items)

	return &items
}

func (s *FileSystemStore) StoreItem(description string) {
	/* s.todo = append(s.todo, model.Item{
		Id: len(s.todo)+1,
		Description: description,
		Order: len(s.todo)+1,
	}) */
}