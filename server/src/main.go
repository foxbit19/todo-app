package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	handler := http.HandlerFunc(TodoServer)
	const port int = 8000
	fmt.Printf("Starting server at port %d\n",port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}