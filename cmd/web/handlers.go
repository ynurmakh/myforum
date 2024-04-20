package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	posts := app.MainModel.GetPosts()
	files := []string{
		"ui/html/base.html",
		"ui/html/partials/nav.html",
		"ui/html/pages/home.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
	}

	err = tmpl.ExecuteTemplate(w, "base", posts)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *Application) postView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *Application) postCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}
