package main

import (
	"fmt"
	"net/http"

	"github.com/GarvitSadhwani/todoApp/controllers"
	"github.com/GarvitSadhwani/todoApp/templates"
	"github.com/GarvitSadhwani/todoApp/views"
	chi "github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml"))
	router.Get("/", controllers.StaticHandler(tpl))
	fmt.Printf("starting server at 8080")
	http.ListenAndServe(":8080", router)

}
