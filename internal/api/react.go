package api

import (
	"html/template"
	"net/http"
)

func (a *api) React(w http.ResponseWriter, r *http.Request) {	
	tmpl := template.Must(template.ParseFiles("frontend/public/views/index.html"))

	tmpl.Execute(w, nil)
}
