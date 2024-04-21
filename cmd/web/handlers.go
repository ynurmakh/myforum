package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Println("not found")
		http.NotFound(w, r)
		return
	}
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
	baseID := path.Base(r.URL.Path)
	id, err := strconv.Atoi(baseID)
	if err != nil || id < 1 {
		fmt.Println("atoi err")
		http.NotFound(w, r)
		return
	}

	post := app.MainModel.GetPost(id)
	files := []string{
		"ui/html/base.html",
		"ui/html/partials/nav.html",
		"ui/html/pages/post-view.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
	}

	err = tmpl.ExecuteTemplate(w, "base", post)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *Application) postCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}
