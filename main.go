package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/GarvitSadhwani/todoApp/controllers"
	"github.com/GarvitSadhwani/todoApp/templates"
	"github.com/GarvitSadhwani/todoApp/views"
	chi "github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	router := chi.NewRouter()
	db, err := sql.Open("pgx", "host=localhost port=5432 user=todoappdb password=todoappdb dbname=simplitask sslmode=disable")
	if err != nil {
		fmt.Printf("error connecting to db")
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("cant comm with db")
	}
	defer db.Close()
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml"))
	router.Get("/", controllers.StaticHandler(tpl))

	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml"))
	router.Get("/signup", controllers.StaticHandler(usersC.Templates.New))

	fmt.Printf("starting server at 8080")
	http.ListenAndServe(":8080", router)

}
