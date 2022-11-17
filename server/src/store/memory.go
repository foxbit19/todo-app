package store

type InMemoryItemStore struct {}

func (s *InMemoryItemStore) GetTodoDescription(id int) string {
	panic("No implementation here")
}

func (s *InMemoryItemStore) StoreItem(description string) {
	panic("No implementation here")
}