package store

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/foxbit19/todo-app/server/src/model"
)

type FileSystemStore struct {
	database io.ReadWriteSeeker
}

func (s *FileSystemStore) GetItem(id int) *model.Item {
	items, _ := s.decodeDatabase()

	for _, item := range *items {
		if(item.Id == id) {
			return &item
		}
	}

	return nil
}

func (s *FileSystemStore) GetItems() *[]model.Item {
	items, _ := s.decodeDatabase()

	return items
}

func (s *FileSystemStore) StoreItem(description string) {
	/* items, _ := s.decodeDatabase()

	items = append(items, model.Item{
		Id: len(items)+1,
		Description: description,
		Order: len(items)+1,
	}) */
}

func (s *FileSystemStore) decodeDatabase() (*[]model.Item, error) {
	s.database.Seek(0, 0)
	var items []model.Item

	err := json.NewDecoder(s.database).Decode(&items)

	if err != nil {
		err = fmt.Errorf("Unable to parse JSON response %v", err)
	}

	return &items, err
}

/* func (s *FileSystemStore) encodeDatabase(*[]model.Item) {
	json.NewEncoder(s.database).Encode(items)

	return &items
} */