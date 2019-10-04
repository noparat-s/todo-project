package store

import (
	"time"
	"todo/pkg/types"
)

type mock struct {
	todos map[int]*types.Todo
}

// NewMock return mock store
func NewMock() Store {
	todos := make(map[int]*types.Todo)
	todos[1] = &types.Todo{
		ID:          1,
		Name:        "Shower",
		CreatedTime: time.Now(),
		DueTime:     time.Now().Add(time.Hour * 48),
	}
	return &mock{
		todos: todos,
	}
}
func (m *mock) GetTodo(id int) *types.Todo {
	return m.todos[id]
}
func (m *mock) Fetch() []*types.Todo {
	return nil
}
func (m *mock) Create(name string, createdTime time.Time, dueTime time.Time) *types.Todo {
	return nil
}
func (m *mock) Update(id int, name string, duetime time.Time) *types.Todo {
	return nil
}
