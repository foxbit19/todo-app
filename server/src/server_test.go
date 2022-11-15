package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETItem(t *testing.T) {
	t.Run("returns the first todo item", func (t *testing.T) {
		// Creates a new GET request on "items" without a body (nil)
		request, _ := http.NewRequest(http.MethodGet, "items", nil)
		// this is the spy
		response := httptest.NewRecorder()

		TodoServer(response, request)

		got := response.Body.String()
		want := "first todo"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}