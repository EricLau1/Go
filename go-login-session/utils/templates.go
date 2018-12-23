package utils

import(
	"net/http"
	"html/template"
)

var templates *template.Template

func LoadTemplates(dir string) {

	templates = template.Must(template.ParseGlob(dir))

}

func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {

	templates.ExecuteTemplate(w, tmpl, data)

}