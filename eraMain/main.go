package main

import (
	"forum/internal/business"
	businessrealiz "forum/internal/business/businessRealiz"
	"forum/internal/storage/sqlite3"
	"log"
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
	test(service)

	// var post &models.Post{

	// }

	// service.CreatePost()
}

func test(service business.Business) {
	_, err := service.GetPostsForHome(1, 20, []int{})
	if err != nil {
		panic(err)
	}

	_, err = service.GetPostByID(1)
	if err != nil {
		panic(err)
	}

	_, err = service.GetPostByID(2)
	if err != nil {
		panic(err)
	}

	_, err = service.GetPostByID(3)
	if err != nil {
		panic(err)
	}
}
