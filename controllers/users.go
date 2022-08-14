package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"sort"
)

type User struct {
	Id        int
	Templates struct {
		New Template
	}
	MyTasks []Task
}

type Data struct {
	Uname string
	Tasks []Task
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

func (u User) HomeHandler(tpl Template, CurrentUserID int, CurrentUserName string) http.HandlerFunc {
	MyTasks := []Task{}
	db, err := sql.Open("pgx", "host=localhost port=5432 user=todoappdb password=todoappdb dbname=simplitask sslmode=disable")
	if err != nil {
		fmt.Println("error connecting to database")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("cant communicate with database")
	}
	res, err := db.Query("select * from tasks where id=$1", CurrentUserID)
	if err != nil {
		fmt.Println("error running query")
	}
	defer res.Close()
	var id int
	var newT string
	var newTD string
	var st int
	var en int
	for res.Next() {
		err = res.Scan(&id, &newT, &newTD, &st, &en)
		if err != nil {
			fmt.Println("error retrieving data from row")
		}
		newTask := Task{
			Task:      newT,
			Detail:    newTD,
			TimeStart: st,
			TimeEnd:   en,
		}
		MyTasks = append(MyTasks, newTask)
	}

	sort.SliceStable(MyTasks, func(i, j int) bool {
		if MyTasks[i].TimeStart == MyTasks[j].TimeStart {
			return MyTasks[i].TimeEnd < MyTasks[j].TimeEnd
		}
		return MyTasks[i].TimeStart < MyTasks[j].TimeStart
	})

	defer db.Close()
	Dataset := Data{
		Uname: CurrentUserName,
		Tasks: MyTasks,
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, Dataset)
	}

}
