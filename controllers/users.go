package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
)

type User struct {
	Id        int
	Templates struct {
		New Template
	}
	MyTasks []Task
}

func (u User) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u User) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "email: ", r.FormValue("email"))
	fmt.Fprint(w, "password: ", r.FormValue("password"))
}

func (u User) HomeHandler(tpl Template, newT string, newTD string) http.HandlerFunc {
	MyTasks := []Task{}
	db, err := sql.Open("pgx", "host=localhost port=5432 user=todoappdb password=todoappdb dbname=simplitask sslmode=disable")
	if err != nil {
		fmt.Println("error connecting to database")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("cant communicate with database")
	}
	res, err := db.Query("select * from tasks where id=1")
	if err != nil {
		fmt.Println("error running query")
	}
	defer res.Close()
	var id int
	for res.Next() {
		err = res.Scan(&id, &newT, &newTD)
		if err != nil {
			fmt.Println("error retrieving data from row")
		}
		newTask := Task{
			Task:   newT,
			Detail: newTD,
		}
		MyTasks = append(MyTasks, newTask)
	}

	defer db.Close()

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, MyTasks)
	}

}

func (u User) Addtask(w http.ResponseWriter, r *http.Request) {
	newTask := Task{
		Task:   r.FormValue("task"),
		Detail: r.FormValue("detail"),
	}
	_ = newTask
	//u.MyTasks = append(u.MyTasks, newTask)

}
