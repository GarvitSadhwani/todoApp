package views

import (
	"fmt"
	"html/template"
	"net/http"
)

type Template struct {
	HTMLTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.HTMLTpl.Execute(w, data)
	if err != nil {
		fmt.Printf("error executing")
	}
}
