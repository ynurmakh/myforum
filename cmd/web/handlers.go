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
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
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
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
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
	if r.Method == http.MethodGet {
		files := []string{
			"ui/html/base.html",
			"ui/html/partials/nav.html",
			"ui/html/pages/post-create.html",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err)
		}

		err = tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		title := r.PostForm.Get("title")
		content := r.PostForm.Get("content")
		id, err := app.MainModel.CreatePost(1, 1, title, content)
		http.Redirect(w, r, fmt.Sprintf("/post/view/%d", id), http.StatusSeeOther)

	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
