package store

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/foxbit19/todo-app/server/model"
	"github.com/foxbit19/todo-app/server/store/utils"
)

// A FileSystemStore stores todo item into a JSON file.
type FileSystemStore struct {
	Database *json.Encoder
	items    []model.Item
}

func NewFileSystemStore(database *os.File) (*FileSystemStore, error) {
	err := initDatabase(database)

	if err != nil {
		return nil, err
	}

	items, err := decodeDatabase(database)

	if err != nil {
		return nil, err
	}

	return &FileSystemStore{
		Database: json.NewEncoder(&utils.Tape{
			File: database,
		}),
		items: items,
	}, nil
}

// GetItem gets an item of the database using its id.
// It decodes the database and call findItem function.
// It returns a reference of the item found, if any, otherwise
// it returns nil.
func (s *FileSystemStore) GetItem(id int) *model.Item {
	return s.findItem(id)
}

// GetItems gets all the items into database.
// It returns a pointer to the slice of the items.
func (s *FileSystemStore) GetItems() *[]model.Item {
	items := s.items
	sort.Slice(items, func (i int, j int) bool  {
		return items[i].Order < items[j].Order
	})

	return &items
}

// StoreItem implements the storing mechanism of an item
// using a description as argument.
// Id is equals to the number of items + 1
func (s *FileSystemStore) StoreItem(description string, order int) int {
	id := len(s.items) + 1
	s.items = append(s.items, model.Item{
		Id:          len(s.items) + 1,
		Description: description,
		Order:       order,
	})

	encodeDatabase(&s.items, s.Database)

	return id
}

// UpdateItem updates an item using the model provided.
// The provided item couldn't be exists.
// The model is used as standard way to pass argument to this function.
// It returns an error if the item cannot be updated
func (s *FileSystemStore) UpdateItem(id int, item *model.Item) error {
	found := s.findItem(id)

	if found == nil {
		return fmt.Errorf("Item %d not found", id)
	}

	if len(item.Description) == 0 {
		return fmt.Errorf("Cannot update item %d without a description", id)
	}

	found.Description = item.Description
	found.Order = item.Order

	encodeDatabase(&s.items, s.Database)

	return nil
}

// DeleteItem deletes an item from this store.
// It uses only item id to looking for the item
// and to delete it.
func (s *FileSystemStore) DeleteItem(id int) {
	index := s.findItemIndex(id)

	// if an element does not exist, nothing occurs
	if index == -1 {
		return
	}

	s.items = append(s.items[:index],s.items[index+1:]...)

	encodeDatabase(&s.items, s.Database)
}

// findItem is a private function to find an item
// inside an array of items.
// The id of the item to find is used to compare items
// between them.
// It returns a reference of the item found, if any, otherwise
// it returns nil.
func (s *FileSystemStore) findItem(id int) *model.Item {
	for i := 0; i < len(s.items); i++ {
		if s.items[i].Id == id {
			return &s.items[i]
		}
	}

	return nil
}

// findItemIndex is a private function to find an index of an item
// inside an array of items.
// The id of the item to find is used to compare items
// between them.
// It returns the index of the item found, if any, otherwise
// it returns nil.
func (s *FileSystemStore) findItemIndex(id int) int {
	for i := 0; i < len(s.items); i++ {
		if s.items[i].Id == id {
			return i
		}
	}

	return -1
}

// decodeDatabase is a private function to decode the database of
// this filesystemStore.
// It uses a json decoder to decode the database and unmarshall it
// to a concrete object (a slice of model.Item).
// It returns the slice and an error (if any).
func decodeDatabase(database io.Reader) ([]model.Item, error) {
	var items []model.Item

	err := json.NewDecoder(database).Decode(&items)

	if err != nil {
		err = fmt.Errorf("Unable to parse JSON response %v", err)
	}

	return items, err
}

// encodeDatabase is a private function to encode the database of
// this filesystemStore.
// It uses the provided json encoder to encode the concrete object and marshall it
// to a Writer.
func encodeDatabase(items *[]model.Item, database *json.Encoder) {
	database.Encode(items)
}

// initDatabase initializes the database if necessary
func initDatabase(database *os.File) (error) {
	database.Seek(0, 0)
	info, err := database.Stat()

	if err != nil {
		return err
	}

	// if the size of the file is equal to 0
	// this piece of code write an empty json array
	// into the file to initialize database
	if info.Size() == 0 {
		database.Write([]byte("[]"))
		database.Seek(0, 0)
	}

	return nil
}