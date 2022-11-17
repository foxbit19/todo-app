package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/foxbit19/todo-app/server/src/model"
)

type StubItemStore struct {
	todo []model.Item
}

func (s *StubItemStore) GetItem(id int) *model.Item {
	for i := 0; i < len(s.todo); i++ {
		if(s.todo[i].Id == id) {
			return &s.todo[i]
		}
	}

	return nil
}

func (s *StubItemStore) GetItems() *[]model.Item {
	return &s.todo
}

func (s *StubItemStore) StoreItem(description string) {
	s.todo = append(s.todo, model.Item{
		Id: len(s.todo)+1,
		Description: description,
		Order: 0,
	})
}

func newGetTodoRequest(id int) *http.Request {
	// Creates a new GET request on "items" without a body (nil)
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/items/%d", id), nil)
	return request
}

func newGetAllTodosRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/items/", nil)
	return request
}

func newPostTodoRequest(item *model.Item) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, "/items/", strings.NewReader(item.Description))
	return request
}

func assertResponseBody(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertResponseStatus(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Not the correct response status: got %d, want %d", got, want)
	}
}

func assertAndGetJsonResponse(t *testing.T, b *bytes.Buffer) *model.Item {
	t.Helper()
	var got model.Item
	err := json.NewDecoder(b).Decode(&got)

	if err != nil {
		t.Errorf("Unable to parse JSON response %q: %v", b, err)
	}

	return &got;
}

func assertAndGetAllJsonResponse(t *testing.T, b *bytes.Buffer) *[]model.Item {
	t.Helper()
	var got []model.Item
	err := json.NewDecoder(b).Decode(&got)

	if err != nil {
		t.Errorf("Unable to parse JSON response %q: %v", b, err)
	}

	return &got;
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

func TestGETTodoItem(t *testing.T) {
	store := StubItemStore{
		[]model.Item{
			{
				Id: 1,
				Description: "this is my first todo",
				Order: 1,
			},
			{
				Id: 2,
				Description: "this is my second todo",
				Order: 2,
			},
		},
	}
	server := NewTodoServer(&store)

	t.Run("returns the first todo item", func(t *testing.T) {
		request := newGetTodoRequest(1)
		// this is the spy
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, "application/json")
		got := assertAndGetJsonResponse(t, response.Body)
		assertResponseBody(t, got.Description, "this is my first todo")
	})

	t.Run("returns the second todo item", func(t *testing.T) {
		request := newGetTodoRequest(2)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, "application/json")
		got := assertAndGetJsonResponse(t, response.Body)
		assertResponseBody(t, got.Description, "this is my second todo")
	})

	t.Run("returns 404 on missing item", func(t *testing.T) {
		request := newGetTodoRequest(0)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusNotFound)
	})

	t.Run("Returns all todo items as JSON array", func(t *testing.T) {
		request := newGetAllTodosRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, "application/json")
		got := assertAndGetAllJsonResponse(t, response.Body)

		if !reflect.DeepEqual(*got, store.todo) {
			t.Errorf("got %q, want %q", *got, store.todo)
		}
	})
}

func TestStoreTodoItems(t *testing.T) {
	store := StubItemStore{
		[]model.Item{},
	}
	server := NewTodoServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/items/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusAccepted)
	})

	t.Run("it stores a todo using POST", func(t *testing.T) {
		item := &model.Item{
			Description: "new todo item",
		}

		request := newPostTodoRequest(item)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusAccepted)

		// verify if the item was correctly stored
		got := server.store.GetItem(1)
		want := "new todo item"

		if got.Description != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
