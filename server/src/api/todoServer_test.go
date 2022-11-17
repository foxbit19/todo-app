package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/foxbit19/todo-app/server/src/model"
)

type StubItemStore struct {
	todo map[int]string
}

func (s *StubItemStore) GetTodoDescription(id int) string {
	return s.todo[id]
}

func (s *StubItemStore) StoreItem(description string) {
	s.todo[len(s.todo)+1] = description
}

func newGetTodoRequest(id int) *http.Request {
	// Creates a new GET request on "items" without a body (nil)
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/items/%d", id), nil)
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

func TestGETTodoItem(t *testing.T) {
	store := StubItemStore{
		map[int]string{
			1: "this is my first todo",
			2: "this is my second todo",
		},
	}
	server := NewTodoServer(&store)

	t.Run("returns the first todo item", func(t *testing.T) {
		request := newGetTodoRequest(1)
		// this is the spy
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "this is my first todo")

	})

	t.Run("returns the second todo item", func(t *testing.T) {
		request := newGetTodoRequest(2)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "this is my second todo")
	})

	t.Run("returns 404 on missing item", func(t *testing.T) {
		request := newGetTodoRequest(0)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreTodoItems(t *testing.T) {
	store := StubItemStore{
		map[int]string{},
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
		got := server.store.GetTodoDescription(1)
		want := "new todo item"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
