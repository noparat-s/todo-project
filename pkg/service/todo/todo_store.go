package todo

import "time"

// Store store interface
type Store interface {
	GetTodo(int) *Todo
	Fetch() *[]Todo
	Create(string, time.Time, time.Time) *Todo
	Update(int, string, time.Time) *Todo
}
