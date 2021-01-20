package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// LoadTemplates - Load all html templates
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// RunTemplate - Run recieved templates
func RunTemplate(w http.ResponseWriter, template string, datas interface{}) {
	templates.ExecuteTemplate(w, template, datas)
}
