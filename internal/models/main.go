package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

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

func openCreate() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	create(db)
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

func insert(db *sql.DB) {
	insertSQL := `
	INSERT INTO users (name, age) VALUES (?, ?)
	`
	_, err := db.Exec(insertSQL, "John Doe", 25)
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
