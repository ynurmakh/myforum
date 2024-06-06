package http

import (
	"fmt"
	"net/http"
)

func (app *App) CookiesMiddlware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allCookies := r.Cookies()
		sessionCookie := &http.Cookie{}

		for i := 0; i < len(allCookies); i++ {
			if allCookies[i].Name == "session" {
				sessionCookie = allCookies[i]
				break
			}
		}

		fmt.Println(sessionCookie.MaxAge)
		fmt.Println(app.Configs.CookiesMaxAge)

		if sessionCookie.Name == "session" {
			sessionCookie.MaxAge = app.Configs.CookiesMaxAge
			http.SetCookie(w, sessionCookie)
			fmt.Println("Update old cookie")
			return
		}

		if sessionCookie.Name != "session" {

			http.SetCookie(w, &http.Cookie{
				Name:   "session",
				Value:  "erbol",
				MaxAge: app.Configs.CookiesMaxAge,
			})

			fmt.Println("Set new cookie")

			return
		}

		next(w, r)
	}
}
