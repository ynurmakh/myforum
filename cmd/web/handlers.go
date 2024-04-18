package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("ui/html/base.html")
	if err != nil {
		log.Println(err)
	}

	tmpl.ExecuteTemplate(w, "base", "forum")
}
