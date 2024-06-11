package http

import (
	"net/http"
)

func (t *Transport) CookiesMiddlware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// sessionCookie, err := r.Cookie("session")
		// if err != nil {

		// 	fmt.Println("ERRR:>>>>>>>>>", err)
		// 	// todo create new cookie for this user

		// 	uuid, err := t.service.GetNewCookie()
		// 	if err != nil {
		// 		// internalerr
		// 		fmt.Println("internal error")
		// 	}

		// 	http.SetCookie(w, &http.Cookie{
		// 		Name:   "session",
		// 		Value:  uuid,
		// 		MaxAge: 5,
		// 	})

		// 	t.User = nil
		// 	next(w, r)
		// 	return
		// }

		// t.User, err = t.service.GetUserByCookiesValues(sessionCookie.Value)
		// if err != nil {
		// 	// todo internal err
		// }

		// if t.User != nil {
		// 	// user in
		// 	fmt.Println(t.User)
		// } else {
		// 	// user NOT
		// }

		// cookie := &http.Cookie{
		// 	Name:   "session",
		// 	Value:  "123",
		// 	MaxAge: 60,
		// }

		// http.SetCookie(w, cookie)
		// cookie, err = r.Cookie("session")
		// if err != nil {
		// 	switch {
		// 	case errors.Is(err, http.ErrNoCookie):
		// 		http.Error(w, "cookie not found", http.StatusBadRequest)
		// 	default:
		// 		log.Println(err)
		// 		http.Error(w, "server error", http.StatusInternalServerError)
		// 	}
		// 	return
		// }

		next(w, r)

		// allCookies := r.Cookies()

		// fmt.Println(allCookies)

		// var sessionCookie *http.Cookie
		// for i := 0; i < len(allCookies); i++ {
		// 	if allCookies[i].Name == "session" {
		// 		sessionCookie = allCookies[i]
		// 		break
		// 	}
		// }

		// // fmt.Println(sessionCookie.MaxAge)
		// // fmt.Println(t.configs.CookiesMaxAge)

		// if sessionCookie != nil {
		// 	if sessionCookie.Name == "session" {
		// 		// sessionCookie.MaxAge = t.configs.CookiesMaxAge
		// 		http.SetCookie(w, sessionCookie)
		// 		fmt.Println("Update old cookie")
		// 	}

		// 	if sessionCookie.Name != "session" {

		// 		http.SetCookie(w, &http.Cookie{
		// 			Name:   "session",
		// 			Value:  "erbol",
		// 			MaxAge: t.configs.CookiesMaxAge,
		// 		})

		// 		fmt.Println("Set new cookie")

		// 	}
		// }
	}

	/*

		cookie 123 >> erbol

	*/
}
