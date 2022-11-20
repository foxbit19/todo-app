package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/foxbit19/todo-app/server/api"
	"github.com/foxbit19/todo-app/server/store"
	"github.com/joho/godotenv"
)

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
  	}
}

func main()  {
	loadEnvFile()
	database, err := os.OpenFile(fmt.Sprintf("./db/%s", os.Getenv("DB_FILE_NAME")), os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("Error on create or open database %s: %v", os.Getenv("DB_FILE_NAME"), err)
	}

	store, err := store.NewFileSystemStore(database)

	if err != nil {
		log.Fatalf("Unable to open database file %s: %v", os.Getenv("DB_FILE_NAME"), err)
	}

	server := api.NewTodoServer(store)

	log.Printf("Starting ToDo server on port %s", os.Getenv("PORT"))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), server); err != nil {
		log.Fatalf("Could not listen on port %s %v", os.Getenv("PORT"), err)
	}

}