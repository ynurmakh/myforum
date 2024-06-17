package main

import (
	"fmt"
	"log"
	"os"

	"forum/internal/models"
	"forum/internal/storage/sqlite3"

	businessrealiz "forum/internal/business/businessRealiz"
)

func main() {
	fmt.Println(len("123"))
	fmt.Println(len("ф"))
	fmt.Println(len("ы"))
	fmt.Println(len("в"))
	fmt.Println(len("汉字汉字汉字汉字汉字汉字汉字汉字汉字汉字"))
	fmt.Println(len("字"))

	os.Exit(1)
	storage, err := sqlite3.InitStorage()
	if err != nil {
		log.Fatal(err)
	}

	service, err := businessrealiz.InitService(storage)
	if err != nil {
		log.Fatal(err)
	}

	bservice := service.(*businessrealiz.Service)
	p, _ := bservice.GetOnlyMyPosts(&models.User{User_id: 1})
	fmt.Println(len(*p))
}
