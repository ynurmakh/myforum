package main

import (
	"database/sql"
	"flag"
	"forum/internal/storage"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Application struct {
	MainModel     storage.MainModel
	TemplateCache map[string]*template.Template
	UserId        int
}

func main() {
	var port string
	flag.StringVar(&port, "p", "8080", "port")
	flag.Parse()

	templateCache, err := newTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app := &Application{
		MainModel:     storage.MainModel{DB: openDB()},
		TemplateCache: templateCache,
	}

	log.Println("server started on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, app.routes()))
}

func openDB() *sql.DB {
	db, err := sql.Open("sqlite3", "internal/storage/database.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
