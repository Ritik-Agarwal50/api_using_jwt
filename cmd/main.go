package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ritik-agarwal50/go_api_tutorial/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterTodoRoutes(r)
	http.Handle("/", r)
	port := 6543
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
