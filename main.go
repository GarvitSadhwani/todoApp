package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/GarvitSadhwani/todoApp/views"
	chi "github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		fmt.Printf("error parsing")
	}
	viewTpl := views.Template{
		HTMLTpl: tpl,
	}
	viewTpl.Execute(w, nil)

}

func main() {
	router := chi.NewRouter()
	router.Get("/", homeHandler)
	fmt.Printf("starting server at 8080")
	http.ListenAndServe(":8080", router)

}
