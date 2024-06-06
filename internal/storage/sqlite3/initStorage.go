package sqlite3

import (
	"database/sql"

	"forum/internal/storage"
)

type Sqlite struct {
	db      *sql.DB
	configs interface{}
}

func InitStorage() (storage.StorageInterface, error) {
	mystorage := &Sqlite{
		configs: "configs here",
	}

	var i storage.StorageInterface
	_ = i
	i = mystorage

	return mystorage, nil
}
