package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/foxbit19/todo-app/server/src/api/common"
	"github.com/foxbit19/todo-app/server/src/model"
	"github.com/foxbit19/todo-app/server/src/store"
	"github.com/gorilla/mux"
)

type TodoServer struct {
	store store.ItemStore
	http.Handler
}

func NewTodoServer(store store.ItemStore) *TodoServer {
	s := &TodoServer{
		store: store,
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", s.welcomeHandler).Methods("GET")
	router.HandleFunc("/items/", s.showItems).Methods("GET")
	router.HandleFunc("/items/{id}", s.showItem).Methods("GET")
	router.HandleFunc("/items/", s.storeItem).Methods("POST")
	router.HandleFunc("/items/{id}", s.updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", s.deleteItem).Methods("DELETE")
	s.Handler = router

	return s
}

// welcomeHander shows a simple welcome message
func (s *TodoServer) welcomeHandler(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Great scott! Welcome to ToDo server!"))
}

// showItems returns all the todo items stored into the store
func (s *TodoServer) showItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(s.store.GetItems())
}

// showItem returns the item that match id argument
func (s *TodoServer) showItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 16)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if id == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(s.store.GetItem(int(id)))
}

// storeItem stores the item into store
func (s *TodoServer) storeItem(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)

	if r.ContentLength == 0 {
		return
	}

	var jsonBody map[string]string
	json.NewDecoder(r.Body).Decode(&jsonBody)
	s.store.StoreItem(jsonBody["description"])
}

// updateItem updates an item using the given id to find it
// and the PUT body to update fields
func (s *TodoServer) updateItem(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 16)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if r.ContentLength == 0 {
		common.ErrorResponse(w, http.StatusBadRequest, "No payload provided")
		return
	}

	var jsonBody map[string]string
	err = json.NewDecoder(r.Body).Decode(&jsonBody)

	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Error on decoding update payload: %v", err))
		return
	}

	err = s.store.UpdateItem(int(id), &model.Item{
		Description: jsonBody["description"],
	})

	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Error on update: %v", err))
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// deleteItem deletes an item from store using only id of the item
func (s *TodoServer) deleteItem(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 16)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	s.store.DeleteItem(int(id))
	w.WriteHeader(http.StatusOK)
}