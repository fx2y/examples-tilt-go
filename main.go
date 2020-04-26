package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", NewExampleRouter())
	log.Println("Serving on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Server exited with: %v", err)
	}
}

type ExampleRouter struct {
	*mux.Router
}

func NewExampleRouter() *ExampleRouter {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./web"))
	r.PathPrefix("/").Handler(fs)
	return &ExampleRouter{Router: r}
}
