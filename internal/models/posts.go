package models

import "time"

type Post struct {
	Post_ID         int64 `db:"post_id"`
	User            User
	Post_Title      string    `db:"post_title"`
	Post_Content    string    `db:"post_content"`
	Created_Time    time.Time `db:"created_time"`
	Post_Categories []Category
	Reactions       ReactionsType
	Comments        []Comment
}

type Category struct {
	Category_id   int    `db:"category_id"`
	Category_name string `db:"category_name"`
}

type ReactionsType struct {
	Likes           int
	Dislikes        int
	ReactByThisUser int // 1 This user liked this object, -1 disliked
}

type Comment struct {
	Comment_Id int64
	User                User
	Commentraie_Content string
	Commentarie_Date    time.Time
	Reactions           ReactionsType
}

func (r *ReactionsType) IsLike() bool {
	return r.ReactByThisUser > 0
}

func (r *ReactionsType) IsDislike() bool {
	return r.ReactByThisUser < 0
}
