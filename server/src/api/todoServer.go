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

	s.store.StoreItem("new todo item")
	fmt.Fprint(w, s.store.GetItem(1))
}

func (s *TodoServer) showItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/items/"))

	if id == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(s.store.GetItem(id))
}