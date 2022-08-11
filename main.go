package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/GarvitSadhwani/todoApp/controllers"
	"github.com/GarvitSadhwani/todoApp/views"
	chi "github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		fmt.Printf("error parsing")
	}
	router.Get("/", controllers.StaticHandler(tpl))
	fmt.Printf("starting server at 8080")
	http.ListenAndServe(":8080", router)

}
