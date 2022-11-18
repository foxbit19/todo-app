package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/foxbit19/todo-app/server/src/api/common"
	"github.com/foxbit19/todo-app/server/src/model"
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
		case http.MethodPut:
			s.updateItem(w, r)
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

func (s *TodoServer) updateItem(w http.ResponseWriter, r *http.Request)  {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/items/"))

	if r.ContentLength == 0 {
		common.ErrorResponse(w, http.StatusBadRequest, "No payload provided")
		return
	}

	var jsonBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&jsonBody)

	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Error on decoding update payload: %v", err))
		return
	}

	err = s.store.UpdateItem(id, &model.Item{
		Description: jsonBody["description"],
	})

	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Error on update: %v", err))
		return
	}

	w.WriteHeader(http.StatusAccepted)
}