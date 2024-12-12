package handlers

import (
	"net/http"
	"text/template"
)

func GetMarkdowns(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("web/template/index.html"))
	tmpl.Execute(w, nil)
}
