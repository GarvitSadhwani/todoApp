package main

import (
	"fmt"
	"net/http"

	chi "github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>creating from scratch</h1>")
}

func main() {
	router := chi.NewRouter()
	router.Get("/", homeHandler)
	fmt.Printf("starting server at 8080")
	http.ListenAndServe(":8080", router)

}
