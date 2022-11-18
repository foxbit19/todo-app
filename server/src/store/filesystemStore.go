package store

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/foxbit19/todo-app/server/src/model"
	"github.com/foxbit19/todo-app/server/src/store/utils"
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
	return s.findItem(&s.items, id)
}

// GetItems gets all the items into database.
// It returns a pointer to the slice of the items.
func (s *FileSystemStore) GetItems() *[]model.Item {
	return &s.items
}

// StoreItem implements the storing mechanism of an item
// using a description as argument.
// Other item's fields are set as follow:
//   - Id is equals to the number of items + 1
//   - Order, as Id field, is equals to the number of items + 1 by default
func (s *FileSystemStore) StoreItem(description string) {
	s.items = append(s.items, model.Item{
		Id:          len(s.items) + 1,
		Description: description,
		Order:       len(s.items) + 1,
	})

	encodeDatabase(&s.items, s.Database)
}

// findItem is a private function to find an item
// inside an array of items.
// The id of the item to find is used to compare items
// between them.
// It returns a reference of the item found, if any, otherwise
// it returns nil.
func (s *FileSystemStore) findItem(items *[]model.Item, id int) *model.Item {
	for _, item := range *items {
		if item.Id == id {
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