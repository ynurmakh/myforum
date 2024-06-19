package http

import (
	"net/http"
)

func (t *Transport) CookiesMiddlware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session")
		if err != nil {

			// get new cookie
			//
			cookie := &http.Cookie{
				Name:  "auth",
				Value: "kek",
				Path:  "/",
			}
			http.SetCookie(w, cookie)
		}

		if !checkAuth(c.Value) {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func checkAuth(s string) bool {
	return false
}
