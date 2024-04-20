package main

import "net/http"

func (app *Application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/post/view", app.postView)
	mux.HandleFunc("/post/create", app.postCreate)

	return mux
}
