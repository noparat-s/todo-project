package store

import (
	"time"
	"todo/pkg/types"
)

// Store store interface
type Store interface {
	GetTodo(int) *types.Todo
	Fetch() []*types.Todo
	Create(string, time.Time, time.Time) *types.Todo
	Update(int, string, time.Time) *types.Todo
}
