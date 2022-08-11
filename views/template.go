package views

import (
	"fmt"
	"html/template"
	"net/http"
)

type Template struct {
	HTMLTpl *template.Template
}

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		fmt.Printf("Error parsing")
	}
	return Template{
		HTMLTpl: tpl,
	}, nil
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.HTMLTpl.Execute(w, data)
	if err != nil {
		fmt.Printf("error executing")
	}
}
