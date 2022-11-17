package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/foxbit19/todo-app/server/src/model"
	"github.com/foxbit19/todo-app/server/src/store"
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

func newPostTodoRequest(t *testing.T, description string) *http.Request {
	body, buffer := map[string]string{"description": description}, new(bytes.Buffer)
   	err := json.NewEncoder(buffer).Encode(body)
	if err != nil {
		t.Errorf("Unable to encode JSON %q: %v", body, err)
    }

	request, _ := http.NewRequest(http.MethodPost, "/items/", buffer)
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

func assertAndGetAllJsonResponse(t *testing.T, b *bytes.Buffer, todo *[]model.Item) {
	t.Helper()
	var got []model.Item
	err := json.NewDecoder(b).Decode(&got)

	if err != nil {
		t.Errorf("Unable to parse JSON response %q: %v", b, err)
	}

	if !reflect.DeepEqual(got, *todo) {
		t.Errorf("got %q, want %q", got, *todo)
	}
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
		assertAndGetAllJsonResponse(t, response.Body, &store.todo)
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

		request := newPostTodoRequest(t, item.Description)
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

func TestStoreItemsAndRetrieveThem(t *testing.T) {
	server := NewTodoServer(&store.InMemoryItemStore{})

	server.ServeHTTP(httptest.NewRecorder(), newPostTodoRequest(t, "I have to do some things, at first"))
	server.ServeHTTP(httptest.NewRecorder(), newPostTodoRequest(t, "Next, I have other things to do"))
	server.ServeHTTP(httptest.NewRecorder(), newPostTodoRequest(t, "Lastly, this is the last thing I have to do"))

	t.Run("Get items", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetAllTodosRequest())

		assertResponseStatus(t, response.Code, 200)
		assertAndGetAllJsonResponse(t, response.Body, &[]model.Item{
			{
				Id: 1,
				Description: "I have to do some things, at first",
				Order: 1,
			},
			{
				Id: 2,
				Description: "Next, I have other things to do",
				Order: 2,
			},
			{
				Id: 3,
				Description: "Lastly, this is the last thing I have to do",
				Order: 3,
			},
		})
	})
}
