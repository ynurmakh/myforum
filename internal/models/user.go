package models

type User struct {
	User_id       int64  `db:"user_id"`
	User_lvl      int64  `db:"user_lvl"`
	User_email    string `db:"user_email"`
	User_nickname string `db:"user_nickname"`
}
