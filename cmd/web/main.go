package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "8080", "port")
	flag.Parse()

	http.HandleFunc("/", home)

	fmt.Println("forum started on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
