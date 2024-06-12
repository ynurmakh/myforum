package main

import (
	"fmt"
	businessrealiz "forum/internal/business/businessRealiz"
	"forum/internal/storage/sqlite3"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args)

	storage, err := sqlite3.InitStorage()
	if err != nil {
		log.Fatal(err)
	}

	service, err := businessrealiz.InitService(storage)
	if err != nil {
		log.Fatal(err)
	}
	_ = service

	storage.CheckTheCookie("0654b38f-c5b7-4b1c-9a71-bb6c4b485cd7", 5)
}

func CheckCookie() {
}
