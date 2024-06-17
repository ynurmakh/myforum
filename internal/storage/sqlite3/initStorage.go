package sqlite3

import (
	"database/sql"
	"fmt"
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

	// err = mystorage.checkScheme()

	var _ storage.StorageInterface = mystorage
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
	return db, nil
}

type PragmasStruct struct {
	id          int64
	name        string
	type_       string
	notnull     int64
	default_    sql.NullString
	primary_key string
}

func (s *Sqlite) checkScheme() error {
	chemePath := "internal/storage/sqlite3/sqlitecScheme/createTable.sql"
	file, err := os.ReadFile(chemePath)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(string(file))
	if err != nil {
		return err
	}

	struc, err := s.pragmaScan("users")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(len(struc))

	for i := 0; i < len(struc); i++ {
		fmt.Println(struc[i])
	}

	return nil
}

func (s *Sqlite) pragmaScan(tablename string) ([]PragmasStruct, error) {
	res := make([]PragmasStruct, 0)

	rows, err := s.db.Query(fmt.Sprintf("PRAGMA table_info(%v)", tablename))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var res1 PragmasStruct

		err := rows.Scan(&res1.id, &res1.name, &res1.type_, &res1.notnull, &res1.default_, &res1.primary_key)
		if err != nil {
			return nil, err
		}

		res = append(res, res1)
	}

	return res, nil
}
