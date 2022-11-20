package common

import (
	"fmt"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, statusCode int, error string)  {
	fmt.Print(error)
	w.WriteHeader(statusCode)
}