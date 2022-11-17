package store

type ItemStore interface {
	GetTodoDescription(id int) string
	StoreItem(description string)
}