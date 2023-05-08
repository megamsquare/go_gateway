package handler

import (
	"fmt"
	"net/http"
)


func initiate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Starting Server")
}