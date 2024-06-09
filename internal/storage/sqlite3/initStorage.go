package sqlite3

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"forum/internal/storage"
)

type Sqlite struct {
	db      *sql.DB
	configs interface{}
}

func InitStorage() (storage.StorageInterface, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}

	mystorage := &Sqlite{
		configs: "configs here",
		db:      db,
	}

	var i storage.StorageInterface
	_ = i
	i = mystorage

	return mystorage, nil
}

func openDB() (*sql.DB, error) {
	dbPath := "internal/storage/sqlite3/database.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	chemePath := "internal/storage/sqlite3/sqlitecScheme/createTable.sql"
	file, err := os.ReadFile(chemePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(string(file))
	if err != nil {
		return nil, err
	}
	return db, nil
}
