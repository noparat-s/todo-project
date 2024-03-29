package todo

import (
	"time"
	"todo/pkg/types"
)

// Todo service interface
type Todo interface {
	GetTodo(int) *types.Todo
	Fetch() *[]types.Todo
	Create(string, time.Time, time.Time) *types.Todo
	Update(int, string, time.Time) *types.Todo
}
