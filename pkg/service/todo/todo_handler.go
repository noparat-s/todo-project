package todo

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	writer := json.NewEncoder(w)
	writer.Encode(todoSv.store.GetTodo(id))
}
