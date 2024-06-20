package http

import (
	"net/http"
)

func (t *Transport) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", t.CookiesMiddlware(t.home))
	// mux.HandleFunc("/page/", t.CookiesMiddlware(t.homePages))
	mux.HandleFunc("/post/view/", t.CookiesMiddlware(t.postView))
	mux.HandleFunc("/myposts/", t.CookiesMiddlware(t.myPosts))
	mux.HandleFunc("/liked/", t.CookiesMiddlware(t.liked))
	mux.HandleFunc("/post/create", t.CookiesMiddlware(t.postCreate))
	mux.HandleFunc("/user/signup", t.CookiesMiddlware(t.signup))
	mux.HandleFunc("/user/login", t.CookiesMiddlware(t.login))
	mux.HandleFunc("/user/logout", t.CookiesMiddlware(t.logout))

	return mux
}
