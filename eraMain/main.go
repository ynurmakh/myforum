package main

import (
	"fmt"
	"log"
	"time"

	"forum/internal/storage/sqlite3"

	businessrealiz "forum/internal/business/businessRealiz"
)

func main() {
	storage, err := sqlite3.InitStorage()
	if err != nil {
		log.Fatal(err)
	}

	service, err := businessrealiz.InitService(storage)
	if err != nil {
		log.Fatal(err)
	}

	bservice := service.(*businessrealiz.Service)
	_ = bservice
	for {

		fmt.Println(storage.CheckTheCookie("cooki", 5))
		time.Sleep(1 * time.Second)
	}
}
