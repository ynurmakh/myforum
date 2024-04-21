package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

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

func openCreate() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db
}

func create(db *sql.DB) {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER
	)`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

// type User struct {
// 	UserID       int       `db:"user_id"`
// 	Username     string    `db:"username"`
// 	Email        string    `db:"email"`
// 	PasswordHash string    `db:"password_hash"`
// 	CreatedAt    time.Time `db:"created_at"`
// }
// type Post struct {
// 	PostID     int       `db:"post_id"`
// 	UserID     int       `db:"user_id"`
// 	CategoryID int       `db:"category_id"`
// 	Title      string    `db:"title"`
// 	Content    string    `db:"content"`
// 	CreatedAt  time.Time `db:"created_at"`
// }
func insert(db *sql.DB) {
	insertSQL := `
	INSERT INTO posts (user_id, category_id, title, content, created_at) VALUES (1, 1, "some title", "some content", current_timestamp)
	`
	_, err := db.Exec(insertSQL)
	if err != nil {
		log.Fatal(err)
	}
}

func get(db *sql.DB) {
	selectSQL := `
	SELECT id, name, age FROM users
	`
	rows, err := db.Query(selectSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}

func update(db *sql.DB) {
	updateSQL := `
	UPDATE users SET age = ? WHERE id = ?
	`
	_, err := db.Exec(updateSQL, 30, 1)
	if err != nil {
		log.Fatal(err)
	}
}

func delete(db *sql.DB) {
	deleteSQL := `
	DELETE FROM users WHERE id = ?
	`
	_, err := db.Exec(deleteSQL, 1)
	if err != nil {
		log.Fatal(err)
	}
}
