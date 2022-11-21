package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func NewGetTodoRequest(id int) *http.Request {
	// Creates a new GET request on "items" without a body (nil)
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/items/%d", id), nil)
	return request
}

func NewGetAllTodosRequest(completed bool) *http.Request {
	path := "/items/"

	if completed {
		path += "completed/"
	}

	request, _ := http.NewRequest(http.MethodGet, path, nil)
	return request
}

func NewPostTodoRequest(t *testing.T, description string) *http.Request {
	body, buffer := map[string]string{"description": description}, new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(body)
	if err != nil {
		t.Errorf("Unable to encode JSON %q: %v", body, err)
	}

	request, _ := http.NewRequest(http.MethodPost, "/items/", buffer)
	return request
}