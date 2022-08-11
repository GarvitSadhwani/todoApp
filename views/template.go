package views

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	HTMLTpl *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		log.Printf("Error parsing template")
	}
	return t
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	tpl, err := template.ParseFS(fs, patterns...)
	if err != nil {
		log.Printf("Error parsing templates")
		return Template{}, nil
	}
	return Template{
		HTMLTpl: tpl,
	}, nil
}

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("Error parsing templates")
		return Template{}, nil
	}
	return Template{
		HTMLTpl: tpl,
	}, nil
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.HTMLTpl.Execute(w, data)
	if err != nil {
		log.Printf("error executing template")
		return
	}

}
