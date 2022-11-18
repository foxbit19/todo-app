package store

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/foxbit19/todo-app/server/src/model"
)

// A FileSystemStore stores todo item into a JSON file.
type FileSystemStore struct {
	database io.ReadWriteSeeker
}

// GetItem gets an item of the database using its id.
// It decodes the database and call findItem function.
// It returns a reference of the item found, if any, otherwise
// it returns nil.
func (s *FileSystemStore) GetItem(id int) *model.Item {
	items, _ := s.decodeDatabase()
	return s.findItem(&items, id)
}

// GetItems gets all the items into database.
// It returns a pointer to the slice of the items.
func (s *FileSystemStore) GetItems() *[]model.Item {
	items, _ := s.decodeDatabase()
	return &items
}

// StoreItem implements the storing mechanism of an item
// using a description as argument.
// Other item's fields are set as follow:
// 	- Id is equals to the number of items + 1
// 	- Order, as Id field, is equals to the number of items + 1 by default
func (s *FileSystemStore) StoreItem(description string) {
	items, _ := s.decodeDatabase()

	items = append(items, model.Item{
		Id: len(items)+1,
		Description: description,
		Order: len(items)+1,
	})

	s.encodeDatabase(&items)
}

// findItem is a private function to find an item
// inside an array of items.
// The id of the item to find is used to compare items
// between them.
// It returns a reference of the item found, if any, otherwise
// it returns nil.
func (s *FileSystemStore) findItem(items *[]model.Item, id int) *model.Item {
	for _, item := range *items {
		if(item.Id == id) {
			return &item
		}
	}

	return nil
}

// decodeDatabase is a private function to decode the database of
// this filesystemStore.
// It uses a json decoder to decode the database and unmarshall it
// to a concrete object (a slice of model.Item).
// It returns the slice and an error (if any).
func (s *FileSystemStore) decodeDatabase() ([]model.Item, error) {
	s.database.Seek(0, 0)
	var items []model.Item

	err := json.NewDecoder(s.database).Decode(&items)

	if err != nil {
		err = fmt.Errorf("Unable to parse JSON response %v", err)
	}

	return items, err
}

// encodeDatabase is a private function to encode the database of
// this filesystemStore.
// It uses a json encoder to encode the concrete object and marshall it
// to a ReadWriteSeeker.
func (s *FileSystemStore) encodeDatabase(items *[]model.Item) {
	s.database.Seek(0, 0)
	json.NewEncoder(s.database).Encode(items)
}