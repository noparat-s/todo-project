package todo

import "time"

// Todo service interface
type Todo interface {
	GetTodo(int) *Todo
	Fetch() *[]Todo
	Create(string, time.Time, time.Time) *Todo
	Update(int, string, time.Time) *Todo
}
