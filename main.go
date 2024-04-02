package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
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
