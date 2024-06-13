package http

import (
	"net/http"
)

func (t *Transport) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// mux.Handle("/protected", sessionMiddleware(http.HandlerFunc(protectedRoute)))
	mux.HandleFunc("/", t.CookiesMiddlware(t.home))
	mux.HandleFunc("/post/view/", t.CookiesMiddlware(t.postView))
	mux.HandleFunc("/myposts/", t.CookiesMiddlware(t.myPosts))
	mux.HandleFunc("/liked/", t.CookiesMiddlware(t.liked))
	mux.HandleFunc("/post/create", t.CookiesMiddlware(t.postCreate))
	mux.HandleFunc("/user/signup", t.CookiesMiddlware(t.signup))
	mux.HandleFunc("/user/login", t.CookiesMiddlware(t.login))
	mux.HandleFunc("/user/logout", t.CookiesMiddlware(t.logout))

	return mux
}

// // SessionToken представляет собой сессионный токен.
// type SessionToken string

// // SessionKey - ключ для хранения сессионных токенов в контексте.
// const SessionKey = "sessionToken"

// // Middleware функция для обработки сессионных токенов.
// func sessionMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Получаем сессионный токен из запроса (например, из cookie).
// 		token := r.Header.Get("Authorization")
// 		if token == "" {
// 			// Нет сессионного токена, создаем новый и добавляем его в контекст.
// 			token = generateSessionToken()
// 			// Устанавливаем токен в cookie
// 			cookie := &http.Cookie{
// 				Name:    "session_token",
// 				Value:   token,
// 				Path:    "/",
// 				Expires: time.Now().Add(10 * time.Minute),
// 			}
// 			http.SetCookie(w, cookie)
// 		}

// 		// Устанавливаем сессионный токен в контекст.
// 		ctx := context.WithValue(r.Context(), SessionKey, SessionToken(token))

// 		// Передаем запрос следующему обработчику с обновленным контекстом.
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

// // Функция для генерации нового сессионного токена.
// func generateSessionToken() string {
// 	// Замените на свою собственную логику генерации токена (например, с использованием UUID, JWT и т. д.)
// 	return "token"
// }

// // Обработчик для защищенного маршрута.
// func protectedRoute(w http.ResponseWriter, r *http.Request) {
// 	// Получаем сессионный токен из контекста.
// 	token, ok := r.Context().Value(SessionKey).(SessionToken)
// 	if !ok {
// 		http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 		return
// 	}

// 	// Проверяем сессионный токен (например, проверяем срок действия, сверяем с базой данных).
// 	if !isValidToken(string(token)) {
// 		http.Error(w, "Invalid session token", http.StatusUnauthorized)
// 		return
// 	}

// 	// Выполняем защищенные действия
// 	fmt.Fprintf(w, "Welcome, user with token: %s", token)
// }

// // Функция для проверки сессионного токена.
// func isValidToken(token string) bool {
// 	// Замените на свою собственную логику проверки токена.
// 	return true
// }
