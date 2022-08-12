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
	Task   string
	Detail string
}

func FAQ(tpl Template) http.HandlerFunc {
	questions := []QnA{
		{
			Ques: "What is the limits on photos",
			Ans:  "As much as you want",
		},
		{
			Ques: "How do I contact support",
			Ans:  "Contact us at garvit.sadh@gmail.com",
		},
		{
			Ques: "How long did it take to work on this",
			Ans:  "Constant efforts and was done quickly",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
