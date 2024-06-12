package models

import "time"

type Post struct {
	Post_ID         int64 `db:"post_id"`
	User            User
	Post_Title      string    `db:"post_title"`
	Post_Content    string    `db:"post_content"`
	Created_Time    time.Time `db:"created_time"`
	Post_Categories []Categories
	Reactions       *ReactionsType
}

type Categories struct {
	Category_id   int    `db:"category_id"`
	Category_name string `db:"category_name"`
}

type ReactionsType struct {
	Likes           int
	Dislikes        int
	ReactByThisUser int
}
