package models

type User struct {
	User_id       int    `db:"user_id"`
	User_lvl      int    `db:"user_lvl"`
	User_email    string `db:"user_email"`
	User_nickname string `db:"user_nickname"`
}
