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
	var newTask string
	var newTaskDetail string
	newTask = ""
	newTaskDetail = ""
	router := chi.NewRouter()
	db, err := sql.Open("pgx", "host=localhost port=5432 user=todoappdb password=todoappdb dbname=simplitask sslmode=disable")
	if err != nil {
		fmt.Println("error connecting to database")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("cant communicate with database")
	}
	defer db.Close()
	// newTask = "new task"
	// newTaskDetail = "some details"
	userCont := controllers.User{}
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "layout.gohtml"))
	router.Get("/", userCont.HomeHandler(tpl, newTask, newTaskDetail))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "layout.gohtml"))
	router.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "addTask.gohtml", "layout.gohtml"))
	router.Get("/addTask", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "layout.gohtml"))
	router.Get("/faq", controllers.FAQ(tpl))

	userCont.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "layout.gohtml"))
	router.Get("/signup", userCont.New)
	router.Post("/users", userCont.Create)

	tpl = views.Must(views.ParseFS(templates.FS, "notFound.gohtml"))
	router.NotFound(controllers.StaticHandler(tpl))

	fmt.Println("Starting server at port: 8080")
	http.ListenAndServe(":8080", router)

}
