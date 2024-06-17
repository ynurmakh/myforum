package main

import (
	"fmt"
	"log"
	"time"

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

	fmt.Println(bservice.ReactionsToComment(3, &models.User{User_id: 1}, 1))
	time.Sleep(1 * time.Second)

	fmt.Println(bservice.ReactionsToComment(4, &models.User{User_id: 1}, 1))
	time.Sleep(1 * time.Second)

	fmt.Println(bservice.ReactionsToComment(5, &models.User{User_id: 1}, 1))
	time.Sleep(1 * time.Second)

	fmt.Println(bservice.ReactionsToComment(6, &models.User{User_id: 1}, 1))
	time.Sleep(1 * time.Second)

	fmt.Println(bservice.ReactionsToComment(3, &models.User{User_id: 1}, -1))
	time.Sleep(1 * time.Second)

	fmt.Println(bservice.ReactionsToComment(4, &models.User{User_id: 1}, -1))
	time.Sleep(1 * time.Second)

	fmt.Println(bservice.ReactionsToComment(5, &models.User{User_id: 1}, -1))
	time.Sleep(1 * time.Second)

	fmt.Println(bservice.ReactionsToComment(6, &models.User{User_id: 1}, -1))
	time.Sleep(1 * time.Second)

	fmt.Println(bservice.ReactionsToComment(3, &models.User{User_id: 1}, -1))
	time.Sleep(1 * time.Second)

	fmt.Println(bservice.ReactionsToComment(4, &models.User{User_id: 1}, -1))
	time.Sleep(1 * time.Second)

	fmt.Println(bservice.ReactionsToComment(5, &models.User{User_id: 1}, -1))
	time.Sleep(1 * time.Second)

	fmt.Println(bservice.ReactionsToComment(6, &models.User{User_id: 1}, -1))
	time.Sleep(1 * time.Second)
}
