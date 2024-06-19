package http

import (
	"context"
	"fmt"
	"net/http"
)

func (t *Transport) CookiesMiddlware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth")
		if err != nil {

			uuid, err := t.service.CreateNewCookie()
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			http.SetCookie(w, &http.Cookie{Name: "auth", Value: uuid, Path: "/"})

			ctx := context.WithValue(r.Context(), "user", nil)
			next(w, r.WithContext(ctx))
			return
		}
		user, err := t.service.GetUserByCookie(cookie.Value)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		fmt.Println(user)
		ctx := context.WithValue(r.Context(), "user", user)
		next(w, r.WithContext(ctx))
	}
}
