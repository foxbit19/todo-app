package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubItemStore struct {
	todo map[int] string
}

func (s *StubItemStore) GetTodoDescription(id int) string {
	return s.todo[id]
}

func TestGETTodoItem(t *testing.T) {
	store := StubItemStore{
		map[int]string{
			1: "this is my first todo",
			2: "this is my second todo",
		},
	}
	server := &TodoServer{&store}

	t.Run("returns the first todo item", func (t *testing.T) {
		request := newGetTodoRequest(1)
		// this is the spy
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "this is my first todo")

	})
	t.Run("returns the second todo item", func (t *testing.T) {
		request := newGetTodoRequest(2)
		// this is the spy
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "this is my second todo")
	})
}

func newGetTodoRequest(id int) *http.Request {
	// Creates a new GET request on "items" without a body (nil)
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/items/%d", id), nil)
	return request;
}

func assertResponseBody(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}