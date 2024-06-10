package models

import "time"

type Post struct {
	Post_ID      int
	User_ID      int
	Post_Title   string
	Post_Content string
	Created_Time time.Time
}
