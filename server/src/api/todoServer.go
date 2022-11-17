package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/foxbit19/todo-app/server/src/store"
)

type TodoServer struct {
	store store.ItemStore
	http.Handler
}

func NewTodoServer(store store.ItemStore) *TodoServer {
	s := &TodoServer{
		store: store,
	}

	router := http.NewServeMux()
	router.Handle("/items/", http.HandlerFunc(s.todoHandler))

	s.Handler = router

	return s
}

func (s *TodoServer) todoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodPost:
			s.storeItem(w,r)
		case http.MethodGet:
			s.showItem(w,r)
	}
}

func (s *TodoServer) storeItem(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)

	if r.ContentLength == 0 {
		return
	}

	var jsonBody map[string]string
	json.NewDecoder(r.Body).Decode(&jsonBody)
	s.store.StoreItem(jsonBody["description"])
	fmt.Printf("json body %v", s.store.GetItem(1))
}

func (s *TodoServer) showItem(w http.ResponseWriter, r *http.Request) {
	arg := strings.TrimPrefix(r.URL.Path, "/items/")

	if arg == "" {
		w.Header().Set("content-type", "application/json")
		// returns all the items
		json.NewEncoder(w).Encode(s.store.GetItems())
		return
	}

	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/items/"))

	if id == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(s.store.GetItem(id))
}