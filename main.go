package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/GarvitSadhwani/todoApp/controllers"
	"github.com/GarvitSadhwani/todoApp/templates"
	"github.com/GarvitSadhwani/todoApp/views"
	chi "github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var currentUserID int
var currentUserName string
var router *chi.Mux
var userCont controllers.User

func addUser(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=todoappdb password=todoappdb dbname=simplitask sslmode=disable")
	if err != nil {
		fmt.Println("error connecting to database")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("cant communicate with database")
	}

	res, err := db.Query("select count(*) from users")

	if err != nil {
		fmt.Println("error running query")
	}
	defer res.Close()
	var count int
	for res.Next() {
		err = res.Scan(&count)
		if err != nil {
			fmt.Println("error retrieving data from row")
		}
	}
	count++
	_, err = db.Exec(`insert into users values($1,$2,$3,$4,$5);`, count, r.FormValue("first_name"), r.FormValue("last_name"), r.FormValue("email"), r.FormValue("password"))
	if err != nil {
		fmt.Println("error entering into database")
	}

	defer db.Close()
	title := r.URL.Path[len("/adduser"):]
	// http.Redirect(w, r, "/"+title, http.StatusSeeOther)
	currentUserID = count
	currentUserName = r.FormValue("first_name")
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "layout.gohtml"))
	router.Get("/homer", userCont.HomeHandler(tpl, currentUserID, currentUserName))
	router.Get("/home", userCont.HomeHandler(tpl, currentUserID, currentUserName))
	http.Redirect(w, r, "/home"+title, http.StatusSeeOther)
}

func authUser(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=todoappdb password=todoappdb dbname=simplitask sslmode=disable")
	if err != nil {
		fmt.Println("error connecting to database")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("cant communicate with database")
	}

	res, err := db.Query("select * from users where email=$1 and password=$2", r.FormValue("email"), r.FormValue("password"))

	if err != nil {
		fmt.Println("error running query")
	}
	defer res.Close()
	var count int
	var f_name string
	var l_name string
	var em string
	var pass string
	for res.Next() {
		err = res.Scan(&count, &f_name, &l_name, &em, &pass)
		if err != nil {
			fmt.Println("error retrieving data from row")
		}
	}
	defer db.Close()
	title := r.URL.Path[len("/loginuser"):]
	currentUserID = count
	currentUserName = f_name
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "layout.gohtml"))
	router.Get("/homer", userCont.HomeHandler(tpl, currentUserID, currentUserName))
	router.Get("/home", userCont.HomeHandler(tpl, currentUserID, currentUserName))
	http.Redirect(w, r, "/home"+title, http.StatusSeeOther)
}

func logoutUser(w http.ResponseWriter, r *http.Request) {
	currentUserID = 0
	currentUserName = ""
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "layout.gohtml"))
	title := r.URL.Path[len("/signout"):]
	router.Get("/homer", userCont.HomeHandler(tpl, currentUserID, currentUserName))
	router.Get("/home", userCont.HomeHandler(tpl, currentUserID, currentUserName))
	http.Redirect(w, r, "/"+title, http.StatusSeeOther)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=todoappdb password=todoappdb dbname=simplitask sslmode=disable")
	if err != nil {
		fmt.Println("error connecting to database")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("cant communicate with database")
	}

	st, _ := strconv.Atoi(r.FormValue("start"))
	en, _ := strconv.Atoi(r.FormValue("end"))
	title := r.URL.Path[len("/newTask"):]
	_, err = db.Exec(`insert into tasks values($1,$2,$3,$4,$5);`, currentUserID, r.FormValue("task"), r.FormValue("detail"), st, en)
	if err != nil {
		fmt.Println("error running query")
	}

	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "layout.gohtml"))
	router.Get("/homer", userCont.HomeHandler(tpl, currentUserID, currentUserName))
	router.Get("/home", userCont.HomeHandler(tpl, currentUserID, currentUserName))
	http.Redirect(w, r, "/home"+title, http.StatusSeeOther)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=todoappdb password=todoappdb dbname=simplitask sslmode=disable")
	if err != nil {
		fmt.Println("error connecting to database")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("cant communicate with database")
	}

	_, err = db.Exec(`delete from tasks where id=$1 and task=$2 and detail=$3;`, currentUserID, r.FormValue("task"), r.FormValue("detail"))
	if err != nil {
		fmt.Println("error running query")
	}
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "layout.gohtml"))
	title := r.URL.Path[len("/deleteTask"):]
	router.Get("/home", userCont.HomeHandler(tpl, currentUserID, currentUserName))
	http.Redirect(w, r, "/home"+title, http.StatusSeeOther)
}

func main() {
	router = chi.NewRouter()
	db, err := sql.Open("pgx", "host=localhost port=5432 user=todoappdb password=todoappdb dbname=simplitask sslmode=disable")
	if err != nil {
		fmt.Println("error connecting to database")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("cant communicate with database")
	}
	defer db.Close()
	userCont = controllers.User{}
	tpl := views.Must(views.ParseFS(templates.FS, "landing.gohtml", "layout_landing.gohtml"))
	router.Get("/", controllers.StaticHandler(tpl))
	tpl = views.Must(views.ParseFS(templates.FS, "signin.gohtml", "layout_landing.gohtml"))
	router.Get("/signin", controllers.StaticHandler(tpl))
	router.Post("/adduser", addUser)
	router.Post("/signout", logoutUser)
	router.Post("/loginuser", authUser)
	router.Post("/deletetask", deleteTask)

	tpl = views.Must(views.ParseFS(templates.FS, "home.gohtml", "layout.gohtml"))
	router.Get("/home", userCont.HomeHandler(tpl, currentUserID, currentUserName))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "layout.gohtml"))
	router.Get("/contact", controllers.StaticHandler(tpl))
	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "layout_landing.gohtml"))
	router.Get("/contact_l", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "addTask.gohtml", "layout.gohtml"))
	router.Get("/addTask", controllers.StaticHandler(tpl))
	router.Post("/newTask", addTask)

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "layout.gohtml"))
	router.Get("/faq", controllers.FAQ(tpl))
	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "layout_landing.gohtml"))
	router.Get("/faq_l", controllers.FAQ(tpl))

	userCont.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "layout_landing.gohtml"))
	router.Get("/signup", userCont.New)
	router.Post("/users", userCont.Create)

	tpl = views.Must(views.ParseFS(templates.FS, "notFound.gohtml"))
	router.NotFound(controllers.StaticHandler(tpl))

	fmt.Println("Starting server at port: 8080")
	http.ListenAndServe(":8080", router)

}
