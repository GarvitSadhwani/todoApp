package controllers

import (
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

type QnA struct {
	Ques string
	Ans  string
}

type Task struct {
	Task      string
	Detail    string
	TimeStart int
	TimeEnd   int
}

func FAQ(tpl Template) http.HandlerFunc {
	questions := []QnA{
		{
			Ques: "What is the limit on number of tasks",
			Ans:  "As much as you want!",
		},
		{
			Ques: "How does it work",
			Ans:  "Add tasks on your home page and they will stay there for as long as you require. Simply click on the cross to delete a task.",
		},
		{
			Ques: "What if I add a task that is to be done before some other tasks",
			Ans:  "The algorithm will show your tasks according to the time you enter, irrespective of when you enter that task.",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
