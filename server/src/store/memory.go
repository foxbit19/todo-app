package store

type InMemoryItemStore struct {}

func (s *InMemoryItemStore) GetTodoDescription(id int) string {
	return "todo time!"
}