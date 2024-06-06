package main

import "forum/internal"

func main() {
	err := internal.Run()
	if err != nil {
		panic(err)
	}
}
