package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/foxbit19/todo-app/server/src/api"
	"github.com/foxbit19/todo-app/server/src/store"
)

func main()  {
	server := api.NewTodoServer(&store.InMemoryItemStore{})
	const port int = 8000
	fmt.Printf("Starting server at port %d\n",port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), server))
}