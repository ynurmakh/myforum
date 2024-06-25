package http

import (
	"context"
	"net/http"
)

func (t *Transport) CookiesMiddlware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authCookie, err := r.Cookie("auth")
		if err != nil {
			newCookiie, err := t.CreateCookie()
			if err != nil {
				http.Redirect(w, r, "/user/login", http.StatusSeeOther)
				return
			}
			http.SetCookie(w, newCookiie)

			ctx := context.WithValue(r.Context(), "user", nil)
			next(w, r.WithContext(ctx))
			return
		}
		user, err := t.service.GetUserByCookie(authCookie.Value)
		if err != nil {
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next(w, r.WithContext(ctx))
	}
}

func (t *Transport) CreateCookie() (*http.Cookie, error) {
	uuid, err := t.service.CreateNewCookie()
	if err != nil {
		return nil, err
	}
	return &http.Cookie{Name: "auth", Value: uuid, Path: "/", MaxAge: t.configs.CookiesMaxAge}, nil
}
