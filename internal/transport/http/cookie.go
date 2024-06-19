package http

import (
	"fmt"
	"net/http"
)

func (t *Transport) CookiesMiddlware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("auth")
		if err != nil {
			fmt.Fprintf(w, "You need to login first!")
			return
			// cookie := &http.Cookie{
			// 	Name:  "auth",
			// 	Value: "kek",
			// 	Path:  "/",
			// }
			// http.SetCookie(w, cookie)
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
