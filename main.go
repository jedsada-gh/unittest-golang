package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	route "github.com/unittest-golang/handler"
)

var (
	httpPort = os.Getenv("PORT")
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", route.HomeHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+httpPort, r))
}
