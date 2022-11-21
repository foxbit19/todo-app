package model

// Item represents a todo item.
// A todo item could be live or completed
type Item struct {
	Id int
	Description string
	Order int
	Completed bool
	CompletedDate string
}