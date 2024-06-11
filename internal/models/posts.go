package models

import "time"

type Post struct {
	Post_ID         int64     `db:"post_id"`
	User_ID         int       `db:"user_id"`
	Post_Title      string    `db:"post_title"`
	Post_Content    string    `db:"post_content"`
	Created_Time    time.Time `db:"created_time"`
	Post_Categories []Categories
}
