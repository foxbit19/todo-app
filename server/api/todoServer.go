package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/foxbit19/todo-app/server/api/common"
	"github.com/foxbit19/todo-app/server/core"
	"github.com/foxbit19/todo-app/server/model"
	"github.com/foxbit19/todo-app/server/store"
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
	router.HandleFunc("/", s.welcomeHandler).Methods(http.MethodGet)
	router.HandleFunc("/items/", s.showItems).Methods(http.MethodGet)
	router.HandleFunc("/items/{id}", s.showItem).Methods(http.MethodGet)
	router.HandleFunc("/items/", s.storeItem).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/items/{id}", s.updateItem).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc("/items/{id}", s.deleteItem).Methods(http.MethodDelete, http.MethodOptions)
	router.HandleFunc("/items/reorder/{sourceId}/{targetId}", s.reorderItem).Methods(http.MethodPatch, http.MethodOptions)
	router.Use(toDoServerCorsMiddleware(router))
	//router.Use(mux.CORSMethodMiddleware(router))
	s.Handler = router

	return s
}

// Enable cors for all origins.
// NOT for production use
func toDoServerCorsMiddleware(r *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, Authorization")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

			if req.Method == "OPTIONS" {
				w.WriteHeader(204)
				return
			}

			next.ServeHTTP(w, req)
		})
	}
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
	json.NewEncoder(w).Encode(core.NewItem(s.store).Get(int(id)))
}

// storeItem stores the item into store
func (s *TodoServer) storeItem(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)

	if r.ContentLength == 0 {
		return
	}

	var jsonBody map[string]string
	json.NewDecoder(r.Body).Decode(&jsonBody)

	core.NewItem(s.store).Create(jsonBody["description"])
}

// updateItem updates an item using the given id to find it
// and the PUT body to update fields
func (s *TodoServer) updateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 16)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if r.ContentLength == 0 {
		common.ErrorResponse(w, http.StatusBadRequest, "No payload provided")
		return
	}

	var item model.Item
	err = json.NewDecoder(r.Body).Decode(&item)

	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Error on decoding update payload: %v", err))
		return
	}

	err = core.NewItem(s.store).Update(&model.Item{
		Id: int(id),
		Description: item.Description,
		Order: int(item.Order),
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

	core.NewItem(s.store).Delete(int(id))
	w.WriteHeader(http.StatusOK)
}

func (s *TodoServer) reorderItem(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	sourceId, err := strconv.ParseInt(vars["sourceId"], 10, 16)

	if err != nil {
		log.Fatalf("Source id %s is not correct", vars["sourceId"])
		w.WriteHeader(http.StatusBadRequest)
	}

	targetId, err := strconv.ParseInt(vars["targetId"], 10, 16)

	if err != nil {
		log.Fatalf("Source id %s is not correct", vars["targetId"])
		w.WriteHeader(http.StatusBadRequest)
	}

	core.NewItem(s.store).Reorder(int(sourceId), int(targetId))
	w.WriteHeader(http.StatusAccepted)
}