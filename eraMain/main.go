package main

import (
	"fmt"
	"log"

	"forum/internal/models"
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

	fmt.Println(bservice.CreateNewUser(&models.User{User_email: "form main@mail.r", User_nickname: "myick"}, "paswwored"))
}
