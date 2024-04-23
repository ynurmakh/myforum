package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
)

type TemplateData struct {
	Data any
}

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
	data := &TemplateData{
		Data: posts,
	}
	app.render(w, http.StatusOK, "home.html", data)
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
	data := &TemplateData{
		Data: post,
	}
	app.render(w, http.StatusOK, "post-view.html", data)
}

func (app *Application) postCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data := &TemplateData{}
		app.render(w, http.StatusOK, "post-create.html", data)
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

func (app *Application) login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data := &TemplateData{}
		app.render(w, http.StatusOK, "login.html", data)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		email := r.PostForm.Get("email")
		pass := r.PostForm.Get("pass")
		user, err := app.MainModel.Login(email, pass)
		fmt.Println("login user:", user)
		if err != nil {
			fmt.Println(err)
			return
		}
		// app.User = user
		http.Redirect(w, r, fmt.Sprintf("/"), http.StatusSeeOther)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
