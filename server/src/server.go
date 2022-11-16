package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type ItemStore interface {
	GetTodoDescription(id int) string
}

type TodoServer struct {
	store ItemStore
}

func (s *TodoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/items/"))

	fmt.Fprint(w, s.store.GetTodoDescription(id))
}