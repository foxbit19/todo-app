package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/foxbit19/todo-app/server/src/model"
	testingCommon "github.com/foxbit19/todo-app/server/src/testing"
	"gotest.tools/v3/assert"
)

func assertResponseBody(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertAndGetJsonResponse(t *testing.T, b *bytes.Buffer) *model.Item {
	t.Helper()
	var got model.Item
	err := json.NewDecoder(b).Decode(&got)

	if err != nil {
		t.Errorf("Unable to parse JSON response %q: %v", b, err)
	}

	return &got
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
	store := testingCommon.StubItemStore{
		[]model.Item{
			{
				Id:          1,
				Description: "this is my first todo",
				Order:       1,
			},
			{
				Id:          2,
				Description: "this is my second todo",
				Order:       2,
			},
		},
	}
	server := NewTodoServer(&store)

	t.Run("returns the first todo item", func(t *testing.T) {
		request := testingCommon.NewGetTodoRequest(1)
		// this is the spy
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, response.Code, http.StatusOK)
		assertContentType(t, response, "application/json")
		got := assertAndGetJsonResponse(t, response.Body)
		assertResponseBody(t, got.Description, "this is my first todo")
	})

	t.Run("returns the second todo item", func(t *testing.T) {
		request := testingCommon.NewGetTodoRequest(2)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, response.Code, http.StatusOK)
		assertContentType(t, response, "application/json")
		got := assertAndGetJsonResponse(t, response.Body)
		assertResponseBody(t, got.Description, "this is my second todo")
	})

	t.Run("returns 404 on missing item", func(t *testing.T) {
		request := testingCommon.NewGetTodoRequest(0)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, response.Code, http.StatusNotFound)
	})

	t.Run("Returns all todo items as JSON array", func(t *testing.T) {
		request := testingCommon.NewGetAllTodosRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, response.Code, http.StatusOK)
		assertContentType(t, response, "application/json")
		assertAndGetAllJsonResponse(t, response.Body, &store.Todo)
	})
}

func TestStoreTodoItems(t *testing.T) {
	store := testingCommon.StubItemStore{
		[]model.Item{},
	}
	server := NewTodoServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/items/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, response.Code, http.StatusAccepted)
	})

	t.Run("it stores a todo using POST", func(t *testing.T) {
		item := &model.Item{
			Description: "new todo item",
		}

		request := testingCommon.NewPostTodoRequest(t, item.Description)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, response.Code, http.StatusAccepted)

		// verify if the item was correctly stored
		got := server.store.GetItem(1)
		want := "new todo item"

		if got.Description != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
