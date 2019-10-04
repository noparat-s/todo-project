package todo

import (
	"context"
	"net/http"
	"todo/pkg/service/todo/store"

	"github.com/gorilla/mux"
)

type todoService struct {
	store store.Store
}

var (
	todoSv *todoService
)

// NewService create todo service
func NewService(ctx context.Context, r *mux.Router) http.Handler {
	createService(r)
	return r
}

func createService(r *mux.Router) {
	todoSv = &todoService{
		store: store.NewMock(),
	}
	r.Methods(http.MethodGet).Path("/{id}").HandlerFunc(getTodo)
}
