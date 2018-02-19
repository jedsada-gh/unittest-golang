package handler

import (
	"fmt"
	"net/http"
)

// HomeHandler is route home
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hellow World!")
}
