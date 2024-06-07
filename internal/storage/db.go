package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type StorageInterface interface {
	// Cookies

	// Registration
	// Login

	// Posts
	// Reactions
}

type UserLogin interface {
	// Сервис отдает БД email юзера.
	//  Если пароль который отправил сервис являееться "GUARANTEED" то сторона storage его не проверяет.
	// А если пароль не "GUARANTEED", то storage проверит совпадение пароля и вернят *User только если верно, иначе nil, nil
	GetUserByEmailAndPass(email string, hashed_password string) (*User, error)
}

type UserRegistration interface {
	// Service отправляет email и storage возвращает ture если сущ. юзер с такой почтой, или false если нет.
	IsExistByEmail(email string) (bool, error)
	// Service отправляет NickName и storage возвращает ture если сущ. юзер с такой NickName, или false если нет.
	IsExistByNickName(email string) (bool, error)
	// Service отправляет email, NickName и password в storage, после storage возвращает true/false
	InsertNewUserByEmailAndPass(email, nickname, password string) (bool, error)
}

type Cookies interface {
	// Service отправляет storage UUID и userEmail и storage привязывает за UUID данного user`а
	TieCookieToUser(UserID int, UUID string) (bool, error)
	// Service отправляет storage UUID кукиса и storage возврашает какой юзер закреплен за данным UUID. Если UUID не сущ. в БД то создаст его и вернет nil,nil. Данный метод автоматический продливает жизннь UUID до time.Now() + expireTime
	CheckTheCookie(cookie string, expireTime int) (*User, error)
	// Service отправляет storage UUID кукиса и storage открепляет user`а за данным UUID
	KillCookie(UUID string) (bool, error)
}

type Posts interface {
	// Отправляеться userID(Владелец поста), заголовог поста, контент поста и катогорий массивом стрингов, в ответ жду номер post_id который был успешно создан
	InsertNewPost(userID int, postTitle, postContent string, categories []string) (int64, error)
	//
	SelectPostByPostID()
	SelectAllPostsByUserID()
	SelectAllPostsByCategory()
}

// there only interface for db
type MainModel struct {
	DB *sql.DB
}

type User struct {
	UserID       int       `db:"user_id"`
	Username     string    `db:"username"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}

type Post struct {
	PostID     int       `db:"post_id"`
	UserID     int       `db:"user_id"`
	CategoryID int       `db:"category_id"`
	Title      string    `db:"title"`
	Content    string    `db:"content"`
	CreatedAt  time.Time `db:"created_at"`
}

type Category struct {
	CategoryID   int    `db:"category_id"`
	CategoryName string `db:"category_name"`
}

type Comment struct {
	CommentID int       `db:"comment_id"`
	PostID    int       `db:"post_id"`
	UserID    int       `db:"user_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

type Like struct {
	LikeID int `db:"like_id"`
	PostID int `db:"post_id"`
	UserID int `db:"user_id"`
}

// Рализация методов в папке с методами
func (m *MainModel) GetPosts() []Post {
	result := []Post{}

	query := `SELECT * FROM posts`
	rows, err := m.DB.Query(query)
	if err != nil {
		fmt.Println(err)
		return result
	}
	defer rows.Close()

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.PostID, &post.UserID, &post.CategoryID, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return result
		}
		result = append(result, post)
	}
	return result
}

func (m *MainModel) GetPost(id int) Post {
	query := "SELECT * FROM posts WHERE post_id = ?"
	row := m.DB.QueryRow(query, id)

	post := Post{}

	err := row.Scan(&post.PostID, &post.UserID, &post.CategoryID, &post.Title, &post.Content, &post.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return post
	}
	return post
}

func (m *MainModel) CreatePost(user_id, category_id int, title, content string) (int64, error) {
	insertSQL := `INSERT INTO posts (user_id, category_id, title, content, created_at) VALUES (?,?,?,?, current_timestamp)`
	res, err := m.DB.Exec(insertSQL, user_id, category_id, title, content)
	if err != nil {
		fmt.Println(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *MainModel) Login(email, pass string) (int, error) {
	query := "SELECT user_id, password_hash FROM users WHERE email = ?"
	row := m.DB.QueryRow(query, email, pass)

	var id int
	var hashedPassword string

	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return id, err
	}

	if hashedPassword != pass {
		return 0, errors.New("password incorrect")
	}
	return id, nil
}

func (m *MainModel) Exists(id int) (bool, error) {
	query := "SELECT true FROM users WHERE user_id = ?"
	row := m.DB.QueryRow(query, id)

	var exists bool

	err := row.Scan(&exists)
	return exists, err
}

func (m *MainModel) GetUser(id int) (User, error) {
	query := "SELECT username FROM users WHERE user_id = ?"
	row := m.DB.QueryRow(query, id)

	var name string

	user := User{}
	err := row.Scan(&name)
	if err != nil {
		return user, err
	}

	user.Username = name
	return user, nil
}

func (m *MainModel) CreateUser(name, email, pass string) (int64, error) {
	insertSQL := `INSERT INTO users (username, email, password_hash, created_at) VALUES (?,?,?, current_timestamp)`
	res, err := m.DB.Exec(insertSQL, name, email, pass)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
