package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	route "github.com/unittest-golang/handler"
)

var (
	httpPort = os.Getenv("PORT")
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", route.HomeHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":"+httpPort, handlers.LoggingHandler(os.Stdout, r))
}
