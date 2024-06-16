package main

import (
	"fmt"
	"forum/internal/models"
	"forum/internal/storage/sqlite3"
	"log"
	"os"
	"strconv"
	"strings"

	businessrealiz "forum/internal/business/businessRealiz"

	"github.com/google/uuid"
)

func main() {
	for {
		fmt.Println(uuid.New())
	}

	/*

		406da5bd-7f1a-42fc-8739-c4be0369d18f:12345635


	*/

	arr := []int{1, 2, 3, 5}
	arrStr := make([]string, 0, len(arr))

	for _, v := range arr {
		arrStr = append(arrStr, strconv.Itoa(v))
	}

	Str := "[" + strings.Join(arrStr, ", ") + "]"
	fmt.Println("|" + Str + "|")

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
	//test(*bservice)

	// var p1 models.Post
	// p1.Post_Title = "New Post Title By service.CreatePost"
	// p1.Post_Content = "New Post Content By service.CreatePost New Post Content By service.CreatePost New Post Content By service.CreatePost New Post Content By service.CreatePost New Post Content By service.CreatePost New Post Content By service.CreatePost "
	// p1.User.User_id = 228

	// fmt.Println(p1)

	// service.CreatePost(&p1, []int{})

	// fmt.Println(p1)
	
	//fmt.Println(bservice.GetCategiries())

	bservice.GetPostsForHome(1, 20, []int{1, 2, 3}, nil)
	
	//fmt.Println(storage.SelectComentByPostID(5))

	f, err := storage.SelectComentByPostID(5)
	if err != nil {
		panic(err)
	}
	for inc,i := range f{
fmt.Println("привет", inc, i)
	}
	
}

func test(service businessrealiz.Service) {
	posts, err := service.GetPostsForHome(1, 2, []int{1, 2, 3}, &models.User{User_id: 228})
	if err != nil {
		panic(err)
	}
	for _, post := range *posts {
		fmt.Println(post)
	}

	fmt.Println(1)
	posts, err = service.GetPostsForHome(2, 2, []int{1, 2, 3}, nil)
	if err != nil {
		panic(err)
	}
	for _, post := range *posts {
		fmt.Println(post)
	}

	os.Exit(1)
	_, err = service.GetPostsForHome(1, 20, []int{}, nil)
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
