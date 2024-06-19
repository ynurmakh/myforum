package http

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func (t *Transport) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.Handle("/protected", sessionMiddleware(http.HandlerFunc(protectedRoute)))
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

// SessionToken представляет собой сессионный токен.
type SessionToken string

// SessionKey - ключ для хранения сессионных токенов в контексте.
const SessionKey = "sessionToken"

// // Middleware функция для обработки сессионных токенов.
func sessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем сессионный токен из запроса (например, из cookie).
		token := r.Header.Get("Authorization")
		if token == "" {
			// Нет сессионного токена, создаем новый и добавляем его в контекст.
			token = generateSessionToken()
			// Устанавливаем токен в cookie
			cookie := &http.Cookie{
				Name:    "session_token",
				Value:   token,
				Path:    "/",
				Expires: time.Now().Add(10 * time.Minute),
			}
			http.SetCookie(w, cookie)
		}

		// Устанавливаем сессионный токен в контекст.
		ctx := context.WithValue(r.Context(), SessionKey, SessionToken(token))

		// Передаем запрос следующему обработчику с обновленным контекстом.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Функция для генерации нового сессионного токена.
func generateSessionToken() string {
	// Замените на свою собственную логику генерации токена (например, с использованием UUID, JWT и т. д.)
	return "token"
}

// Обработчик для защищенного маршрута.
func protectedRoute(w http.ResponseWriter, r *http.Request) {
	// Получаем сессионный токен из контекста.
	ctx := r.Context()
	ctxValue := ctx.Value(SessionKey)
	token, ok := ctxValue.(SessionToken)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Проверяем сессионный токен (например, проверяем срок действия, сверяем с базой данных).
	if !isValidToken(string(token)) {
		http.Error(w, "Invalid session token", http.StatusUnauthorized)
		return
	}

	// Выполняем защищенные действия
	fmt.Fprintf(w, "Welcome, user with token: %s", token)
}

// Функция для проверки сессионного токена.
func isValidToken(token string) bool {
	// Замените на свою собственную логику проверки токена.
	return true
}

type User struct {
	Username string
	Password string
}

// Mock user database (replace with your actual authentication logic)
var users = map[string]*User{
	"testuser": {Username: "testuser", Password: "password"},
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		user, ok := users[username]
		if ok && user.Password == password {
			// Create a cookie to store the username
			cookie := http.Cookie{
				Name:     "username",
				Value:    username,
				Path:     "/",
				Expires:  time.Now().Add(15 * time.Minute), // Set cookie expiration
				HttpOnly: true,
			}

			http.SetCookie(w, &cookie)
			fmt.Fprintf(w, "Login successful! You are redirected to protected page in 5 seconds")
			// Optionally redirect to protected page after a short delay
			time.Sleep(5 * time.Second)
			http.Redirect(w, r, "/protected", http.StatusFound)
		} else {
			fmt.Fprintf(w, "Invalid username or password")
		}
	} else {
		// Display login form if it's a GET request
		fmt.Fprintf(w, `<html>
		<body>
		<h2>Login</h2>
		<form method="POST" action="/login">
		Username: <input type="text" name="username"><br>
		Password: <input type="password" name="password"><br>
		<input type="submit" value="Login">
		</form>
		</body>
		</html>`)
	}
}

func handleProtected(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintf(w, "You need to login first!")
		return
	}

	// Access the username from the cookie
	username := cookie.Value

	fmt.Fprintf(w, "Welcome, %s! You are logged in.", username)
}

// Root handler for demonstration
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome! You can login at /login")
}

// func main() {
// 	http.HandleFunc("/", handleRoot)
// 	http.HandleFunc("/login", handleLogin)
// 	http.HandleFunc("/protected", handleProtected)
// 	http.ListenAndServe(":8080", nil)
// }
