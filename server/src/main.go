package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/foxbit19/todo-app/server/src/api"
	"github.com/foxbit19/todo-app/server/src/store"
)

const dbFileName = "todo.db.json"
const port = 8000

func main()  {
	database, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("Error on create or open database %s: %v", dbFileName, err)
	}

	store, err := store.NewFileSystemStore(database)

	if err != nil {
		log.Fatalf("Unable to open database file %s: %v", dbFileName, err)
	}

	server := api.NewTodoServer(store)

	log.Printf("Starting ToDo server on port %d",port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), server); err != nil {
		log.Fatalf("Could not listen on port %d %v", port, err)
	}

}