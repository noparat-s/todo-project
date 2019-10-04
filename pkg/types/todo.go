package types

import "time"

// Todo model
type Todo struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	CreatedTime time.Time `json:"createdTime"`
	DueTime     time.Time `json:"dueTime"`
}
