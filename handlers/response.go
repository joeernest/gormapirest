package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// sendData : Send a data with responseWriter and status code
func sendData(rw http.ResponseWriter, data interface{}, status int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)

	output, _ := json.Marshal(&data)
	fmt.Fprintf(rw, string(output))
}

// sendError : Send a responseWriter with status code and error message
func sendError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	fmt.Fprintf(rw, "Resource not found: %d", status)
}
