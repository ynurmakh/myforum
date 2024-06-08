package http

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func (t *Transport) CookiesMiddlware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie := &http.Cookie{
			Name:  "session",
			Value: "KEK",
		}

		http.SetCookie(w, cookie)
		cookie, err := r.Cookie("exampleCookie")
		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				http.Error(w, "cookie not found", http.StatusBadRequest)
			default:
				log.Println(err)
				http.Error(w, "server error", http.StatusInternalServerError)
			}
			return
		}
		fmt.Println(cookie.Value, err)

		// allCookies := r.Cookies()

		// fmt.Println(allCookies)

		// var sessionCookie *http.Cookie
		// for i := 0; i < len(allCookies); i++ {
		// 	if allCookies[i].Name == "session" {
		// 		sessionCookie = allCookies[i]
		// 		break
		// 	}
		// }

		// fmt.Println(sessionCookie.MaxAge)
		// fmt.Println(t.configs.CookiesMaxAge)

		// if sessionCookie.Name == "session" {
		// 	sessionCookie.MaxAge = t.configs.CookiesMaxAge
		// 	http.SetCookie(w, sessionCookie)
		// 	fmt.Println("Update old cookie")
		// }

		// if sessionCookie.Name != "session" {

		// 	http.SetCookie(w, &http.Cookie{
		// 		Name:   "session",
		// 		Value:  "erbol",
		// 		MaxAge: t.configs.CookiesMaxAge,
		// 	})

		// 	fmt.Println("Set new cookie")

		// }

		next(w, r)
	}

	/*

		cookie 123 >> erbol

	*/
}
